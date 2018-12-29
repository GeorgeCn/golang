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
		for x := 0; x < 100; x++ {
			time.Sleep(1 * time.Second)
			naturals<- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for x := range naturals {
			squares<- x * x
		}
	}()

	//printer
	for x := range squares {
		fmt.Println(x)
	}
}
