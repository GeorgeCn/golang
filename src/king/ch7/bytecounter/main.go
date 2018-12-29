package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // 转换 int 为 ByteCounter类型

	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0 //重置计数器
	var name = "Dolly"
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)
}

