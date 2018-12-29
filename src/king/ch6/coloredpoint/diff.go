package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) PrintInfoPointer() {
	fmt.Printf("%p,%v", p, p)
}

func (p Person) PrintInfoValue() {
	fmt.Printf("%p, %v\n", &p, p)
}

func main() {
	p := Person{"乔治", 1, 18}
	p.PrintInfoPointer()
	fmt.Println("---------------\n")

	//方法表达式
	pFunc1 := (*Person).PrintInfoPointer
	pFunc1(&p)
	pFunc2 := Person.PrintInfoValue
	pFunc2(p)
	fmt.Println("---------------\n")

	//方法值
	pFunc3 := p.PrintInfoPointer
	pFunc3()

	pFunc4 := p.PrintInfoValue
	pFunc4()
	fmt.Println("---------------\n")
}