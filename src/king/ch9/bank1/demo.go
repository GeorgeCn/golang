package main

import "fmt"

func main()  {
	demo := make(chan int)
	go func() {
		<-demo
	}()
	demo<- 2

	fmt.Println(<-demo)
}
