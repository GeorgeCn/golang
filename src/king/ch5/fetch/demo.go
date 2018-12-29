package main

import "fmt"

func main() {
	defer fmt.Println(111)
	defer fmt.Println(222)
	defer fmt.Println(333)
	defer fmt.Println(444)
}
