package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
)

// WaitForServer 尝试连接URL对应的服务器
//在一分钟内使用指数退避策略进行重试
//所以的尝试失败后返回错误

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << unit(tries)) // 指数退避策略
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
