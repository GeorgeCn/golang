package main

import "sync"

var(
	mu sync.Mutex // 保护 mapping
	mapping = make(map[string]string)
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string)
}

func main() {

}

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

