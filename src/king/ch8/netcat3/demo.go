package main

import (
	"fmt"
	"time"
)

func main() {
	ch :=make(chan int)

	go func() {
		ch<- 1
		fmt.Println(111)
	}()
	time.Sleep(1*time.Second)
	fmt.Println("2")
	<- ch
	time.Sleep(1*time.Second)
}
