package main

import (
	"sync"
	"fmt"
)

//var cache = struct {
//	sync.Mutex
//	mapping map[string]string
//}{
//	mapping:make(map[string]string),
//}

type demo struct {
	sync.Mutex
	mapping map[string]string
}

func main() {
	cache := demo{{1,1}, mapping:{"hello":"world"}}
	fmt.Println(cache)
}
