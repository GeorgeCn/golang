package main

import "fmt"

type demo interface {
	echo() string
}

type demo_struct struct {
	X int
	Y int
}

func(d demo_struct) echo() string {
	return "Hello World?"
}

func main() {
	var w demo
	r := demo_struct{1, 1}
	w = r
	b, ok := w.(demo_struct)
	fmt.Printf("%#v", b)
	fmt.Println(ok)
	fmt.Println(w.(demo))
}
