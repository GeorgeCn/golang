package main

import "fmt"

func main() {
	b := make([]int, 4, 10)
	d := b[:5]
	e := append(d, 5)
	e[1] = 1
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(b)
	test(b)
}

func test(s []int) {
	s[0] = 1
	fmt.Println(s)
}

