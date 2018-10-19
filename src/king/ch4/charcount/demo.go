package main

import (
	"os"
	"fmt"
)

func main() {
	input, err := os.Stdin.Read("./main.go")
	if(err == nil) {
		fmt.Println(input)
	}
}
