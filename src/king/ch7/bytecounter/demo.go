package main

import "fmt"

type DemoInt int

func (demoint *DemoInt) String() {
	fmt.Println(*demoint)
}

func main()  {
	var a DemoInt = 1
	a.String()
}