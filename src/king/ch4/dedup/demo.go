package main

import "fmt"

var m string

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func main() {
	a := []string{"Hello world"}
	m = k(a)
	fmt.Printf("%v", m)
}