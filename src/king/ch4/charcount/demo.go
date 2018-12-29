package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	//counts := make(map[rune]int)
	//var utflen [utf8.UTFMax + 1]int
	//invalid :=0
	var stringss []byte
	input, err := os.Stdin.Read(stringss[:]);
	if err != nil {
		fmt.Println("error")
	}

	in := bufio.NewReader(strings.NewReader(string(stringss)))
	//fmt.Println(utflen)
	fmt.Println(input)
	fmt.Println(in)
}
