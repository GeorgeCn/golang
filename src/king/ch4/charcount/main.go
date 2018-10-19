package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"unicode"
	"fmt"
)

func main() {
	counts := make(map[rune]int) // Unicode字符数量
	var utflen [utf8.UTFMax +1] int //UTF-8 编码的长度
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()// 返回 rune、nbytes、error
		if err == io.EOF {
			break
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
