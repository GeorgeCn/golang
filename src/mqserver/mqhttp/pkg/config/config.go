package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"mqserver/mqhttp/pkg/utils"
	"os"
)

var jsonData map[string]interface{}

//读取配置文件
func initJson() {
	goPath := os.Getenv("GOPATH")
	bytes, err := ioutil.ReadFile(goPath + "/src/mqserver/application.json")
	//bytes, err := ioutil.ReadFile("./application.json")
	if err != nil {
		log.Panic("ReadFile:", err.Error())
	}
	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)
	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		log.Panic("invalid config:", err.Error())
	}
}


type rabbitMQConfig struct {
	Host string
	Port int
	VirtualHosts string
	Url string
	Username string
	Password string
	//Channel string
	Queue string
	ExChange string

}

var RabbitMQConfig rabbitMQConfig

func initRabbitMQ(){
	err := utils.SetStructByJSON(&RabbitMQConfig, jsonData["rabbitmq"].(map[string]interface{}))
	if err != nil {
		log.Panic("rabbitmq config error:", err.Error())
	}
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",RabbitMQConfig.Username,RabbitMQConfig.Password ,RabbitMQConfig.Host, RabbitMQConfig.Port)
	RabbitMQConfig.Url = url
	fmt.Println(url)
}

func init() {
	initJson()
	initRabbitMQ()
}
