package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,4}
	fmt.Println("a:",len(a), cap(a), a)

	b := [10]int{1,2,3,4}
	fmt.Println("b:",len(b), cap(b), b)

	c := make([]int, 4, 10)
	fmt.Println("c:",len(c), cap(c),c)

	d := b[:5]
	fmt.Println("d:",len(d), cap(d),d)
	e := append(d,5) //append后d的容量不变
	e[0] = 100//没超出底层数组的容量，因此e和d都指向同一个数组，修改e会影响d
	fmt.Println("d after append:",len(d), cap(d),d)
	fmt.Println("e:",len(e), cap(e),e)
}
