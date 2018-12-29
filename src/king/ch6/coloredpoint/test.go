package main

import "fmt"

func main() {
	a()
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance // 方法表达式
	fmt.Println(distance(p, q)) // "5"
	fmt.Printf("%T\n")
}

var a = func () {
	fmt.Println("Hello world")
}


