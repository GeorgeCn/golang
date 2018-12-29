package main

import (
	"net/url"
	"fmt"
)

func main() {
	m := url.Values{"lang" : {"en"}} // 直接构造
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3")
}
