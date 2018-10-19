package main

import (
	"fmt"
)

type Rectangle struct {
	length , width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r1 := Rectangle{4, 3}
	temperature := Celsius(100)
	fmt.Printf("Rectangle is:%v", r1)
	fmt.Printf("Rectangle-area is:%d", r1.Area())
	fmt.Println("Temperature is:%s", temperature.String())
}
