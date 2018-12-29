package main

import "fmt"

type demo struct {
	X int
	Y int
}

func main() {
	a := demo{1,2}
	fmt.Println(a.X)
	fmt.Println(demo.demo)
	fmt.Println(a.demo)
}

func (d demo) demo() {
	fmt.Println(d.X)
}
