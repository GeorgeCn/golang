//Author : dmc
//
//Date: 2018/8/13 下午2:45
//
//Description:
package request

import "fmt"

type SendRequest struct {
	ExchangeOptions map[string]interface{} `json:"exchange_options"`
	QueueName       string                 `json:"queue_name"`
	Message         string                 `json:"message"`
	DelayTime       int                    `json:"delay_time"`
}

func (s *SendRequest) String() string{
	return fmt.Sprintf("%#v",s)
}
