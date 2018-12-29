package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0;; x++ {
			time.Sleep(1 * time.Second)
			naturals<- x
		}
	}()

	//squarer
	go func() {
		for {
			x := <- naturals
			squares<- x * x
		}
	}()

	// printer(在主goroutine中)
	for{
		fmt.Println(<- squares)
	}
}
