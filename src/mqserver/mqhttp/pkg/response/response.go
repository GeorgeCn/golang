//Author : dmc
//
//Date: 2018/8/13 下午2:46
//
//Description:
package response

import "fmt"

type ResultJson struct {
	Errcode int         `json:"errcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}


func NewResultJson(errcode int , msg string , data interface{}) ResultJson{
	return ResultJson{Errcode:errcode,Msg:msg,Data:data}
}

func (r ResultJson) String() string{
	return fmt.Sprintf("%#v",r)
}