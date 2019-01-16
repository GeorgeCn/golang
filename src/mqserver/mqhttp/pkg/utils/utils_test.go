//Author : Dom
//
//Date: 2018/7/13 14:17
//
//Description:
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"testing"
	"time"
)

var jsonData map[string]interface{}

//getJSON 读取配置文件
func getJSON() {
	bytes, err := ioutil.ReadFile("../application.json")
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

//数据库配置(正确)
type dbConfig1 struct {
	Dialect      string //数据库类型
	Database     string //数据库名称
	User         string //用户名
	Password     string //密码
	Charset      string //编码
	Host         string //地址
	Port         int    //端口
	SQLLog       bool   //是否打印sql
	URL          string //程序生成
	MaxIdleConns int    //空闲时最大的连接数
	MaxOpenConns int    //最大的连接数
}

//数据库配置(错误)
type dbConfig2 struct {
	Dialects     string //数据库类型
	Database     string //数据库名称
	User         string //用户名
	Password     string //密码
	Charset      string //编码
	Host         string //地址
	Port         int    //端口
	SQLLog       bool   //是否打印sql
	URL          string //程序生成
	MaxIdleConns int    //空闲时最大的连接数
	MaxOpenConns int    //最大的连接数
}

var DBConfig1 dbConfig1
var DBConfig2 dbConfig2

func TestSetStructByJSON(t *testing.T) {
	getJSON()
	err1 := SetStructByJSON(&DBConfig1, jsonData["database"].(map[string]interface{}))
	err2 := SetStructByJSON(&DBConfig2, jsonData["database"].(map[string]interface{}))
	if err1 != nil {
		t.Errorf("SetStructByJSON(%v),err!", DBConfig1)
	}
	if err2 == nil {
		t.Errorf("SetStructByJSON(%v),err!", DBConfig2)
	}
}

func TestMakeIndexToUTC(t *testing.T) {
	x := MakeIndexToUTC(1838)
	fmt.Println(x)
}

func TestMakeUTCToIndex(t *testing.T) {
	tim, _ := time.Parse("2016-01-02 15:04:05", "2018-01-07 09:05:00")
	minutes := tim.Sub(baseTime).Minutes()
	index := int(minutes)/5 + 1
	fmt.Println(index)
}

func TestParseToUTC(t *testing.T) {
	_,err := ParseToUTC("2018-12-371T00:00:00+08:00")
	if err != nil {
		fmt.Println(err.Error())
	}
}