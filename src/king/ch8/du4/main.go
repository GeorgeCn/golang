package main

import (
	"os"
	"sync"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte,1))// 读取一个字符
		close(done)
	}()

	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {

	}

	for {
		select{
		case <-done:
			//耗尽filesize以允许已有的goroutine结束
			for range fileSize {
				//不执行任何操作
			}
			return
		}
	}
}