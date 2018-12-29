package main

import (
	"time"
	"log"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // 别忘记这对括号
	time.Sleep(10 * time.Second) // 通过休眠仿真慢动作
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
