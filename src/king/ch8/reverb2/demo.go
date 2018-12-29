package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
			fmt.Println("123")
			time.Sleep(time.Second)
		}()
	fmt.Println("Hello World")
	go func() {
			fmt.Println("-------------------------------")
			fmt.Println("321")
			time.Sleep(time.Second)
	}()
	time.Sleep(1000*time.Second)
}
