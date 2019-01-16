//Author : dmc
//
//Date: 2018/8/13 下午7:14
//
//Description:
package service

import (
	"fmt"
	"log"
)


type consumer struct {
	queueName string
	routeKey  string
}


func (c *consumer) QueueName() string {
	return c.queueName
}

func (c *consumer) RouterKey() string {
	return c.routeKey
}

func (c *consumer) OnError(err error) {
	log.Println(err)
}

func (c *consumer) OnReceive(body []byte) bool {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	b := NewBasicServiceImpl()
	var msg = string(body)
	err := b.SendMsgToServer(msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func NewConsumer(queueName, routeKey string) *consumer {
	return &consumer{queueName, routeKey}
}
