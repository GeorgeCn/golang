package service

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"mqserver/mqhttp/pkg/rabbitmq"
	"mqserver/mqhttp/pkg/request"
	"mqserver/mqhttp/pkg/response"
	"strconv"
	"time"
)

var (
	XDelayedMessage = "x-delayed-message"
	XDelay          = "x-delay"
	queues          = make(consClients, 0)
)

//消费者客户端列表
type consClient struct {
	queueName string //队列名称
	expire    int64  //过期时间
	client    *rabbitmq.RabbitMQ
}

type consClients []*consClient

//通过队列名称查找客户端
func (c consClients) FindByQname(qName string) *consClient {
	for _, v := range c {
		if v.queueName == qName {
			return v
		}
	}
	return nil
}

// MqhttpService describes the service.
type MqhttpService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendToMQ(ctx context.Context, msg request.SendRequest) response.ResultJson
	ReceiveFromMQ(ctx context.Context)
}

type basicMqhttpService struct{}

func (b *basicMqhttpService) SendToMQ(ctx context.Context, req request.SendRequest) response.ResultJson {
	var exname, extype string
	var args amqp.Table
	//获取exname
	if v, ok := req.ExchangeOptions["name"]; ok {
		exname = v.(string)
	}
	//获取extype
	if v, ok := req.ExchangeOptions["type"]; ok {
		extype = v.(string)
	}
	qname := req.QueueName
	msg := req.Message
	delayTime := req.DelayTime
	//将arguments转成table
	if v, ok := req.ExchangeOptions["arguments"]; ok {
		args = amqp.Table(v.(map[string]interface{}))
	}
	//获得一个新的mqclient
	mq, err := rabbitmq.NewRabbitMQ(exname, extype, args)
	defer mq.Close()
	if err != nil {
		return response.NewResultJson(110001002, "MQ连接失败，请重试", "")
	}
	/*******************开启消费者监听*********************************/
	consumer := NewConsumer(qname, qname)
	if client := queues.FindByQname(qname); client == nil {
		//设置过期时间
		addTime := 600
		if delayTime > 600 {
			addTime = delayTime + 60
		}
		//创建一个新连接，监听
		fmt.Println("正在创建新的消费者队列。。。")
		consumerClient, _ := rabbitmq.CreateNewRabbitMQ(exname, extype, nil)
		queues = append(queues, &consClient{qname, time.Now().UTC().Unix() + int64(addTime), consumerClient})
		go consumerClient.Listen(consumer)
	} else {
		//设置过期时间
		addTime := 600
		if delayTime > 600 {
			addTime = delayTime + 60
		}
		client.expire = time.Now().UTC().Unix() + int64(addTime)
	}
	/****************************************************/

	//defer mq.Close()
	header := make(map[string]interface{})
	//如果delaytime不为0，就延迟发送
	if delayTime != 0 || extype == XDelayedMessage {
		header[XDelay] = strconv.Itoa(delayTime)
	}
	if delayTime == 0 {

	}
	//封装消息格式
	publishing := amqp.Publishing{
		Headers:     amqp.Table(header),
		ContentType: "text/plain",
		Body:        []byte(msg),
	}
	//发布
	succeed := mq.Publish(exname, qname, publishing)
	if succeed {
		return response.NewResultJson(0, "信息发送到MQ,发送成功", "")
	}
	return response.NewResultJson(110001001, "发送失败，请重试", "")
}
func (b *basicMqhttpService) ReceiveFromMQ(ctx context.Context) {
	// TODO implement the business logic of ReceiveFromMQ
}

func (b *basicMqhttpService) SendMsgToServer(msg string) error {
	return sendMsgToServer(msg)
}

func sendMsgToServer(msg string) error {
	fmt.Println(msg)
	return nil
}

// NewBasicMqhttpService returns a naive, stateless implementation of MqhttpService.
func NewBasicMqhttpService() MqhttpService {
	return &basicMqhttpService{}
}

func NewBasicServiceImpl() *basicMqhttpService {
	return &basicMqhttpService{}
}

// New returns a MqhttpService with all of the expected middleware wired in.
func New(middleware []Middleware) MqhttpService {
	var svc = NewBasicMqhttpService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

//移除过期的consumerclient
func cleanExpiredClient() {
	for {
		time.Sleep(time.Second * 5)
		for i, v := range queues {
			if v.expire < time.Now().UTC().Unix() {
				v.client.Close()
				fmt.Println("有连接过期了，已经被移除", v.queueName)
				if i == len(queues)-1 {
					queues = queues[:i]
				} else {
					queues = append(queues[:i], queues[i+1:]...)
				}
			}
		}
	}
}

func init() {
	go cleanExpiredClient()
}
