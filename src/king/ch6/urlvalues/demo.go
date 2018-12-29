package main

import "fmt"

// IntList是整型链表
// *IntList的类型nil代表空列表
type IntList struct {
	Value int
	Tail *IntList
}

// sum返回列表元素的总和
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

func main()  {
	fmt.Println(string(1))
}