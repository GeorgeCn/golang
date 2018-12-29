package main

import "fmt"

//func double(x int) (result int) {
//	result = x
//	fmt.Println(111)
//	return
//}

func demo(x int) (result int)  {
	defer func() {fmt.Printf("double(%d) = %d\n", x, result)}()
	return x + x
}

func main() {
	//_ = double(10)
	//resp := demo(4)
	//fmt.Println(resp)
	fmt.Println(triple(4))
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return demo(x)
}