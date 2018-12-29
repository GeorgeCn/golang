package main

import (
	"fmt"
	"time"
)

func main() {
	var x = 0
	go func() {
		x += 100
		fmt.Println(x)
	}()

	go func() {
		x += 200
		fmt.Println(x)
	}()
	time.Sleep(1000 * time.Second)
	fmt.Println(123)
}
