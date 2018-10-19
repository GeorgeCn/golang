package main

import "fmt"

var x uint8 = 1 <<1 | 1 <<5
func main()  {
	fmt.Printf("%d\n", x)
}
