package main

import (
	"fmt"
	"os"
	"io"
	"bytes"
)

func main()  {
	var w io.Writer
	fmt.Printf("%T\n", w)// 接口值

	w = os.Stdout
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
}
