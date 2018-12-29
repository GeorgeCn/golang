package main

import "fmt"

func sum(primary int, vals ...int) int {
	fmt.Printf("params: %v", primary)
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main()  {
	fmt.Println(sum(1, 3, 4))
}
