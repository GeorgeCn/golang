package main

import "fmt"

type tree struct {
	label    int
}
func demoStruct(trees tree) *tree {
	trees.label = 2
	return &trees
}

func main() {
	trees := tree{1}
	demo := demoStruct(trees)
	fmt.Println(*demo)

	a := 122
	fmt.Printf("|%-5d|", 1)
	fmt.Printf("|%05d|", a)
	fmt.Printf("|%3.2s|", "Hello world")
}
