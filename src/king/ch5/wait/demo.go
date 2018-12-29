package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.Minute
	url := "www.baidu.com"
	fmt.Println(url)
	fmt.Sprintf("server %s failed to respond after %s", url, timeout)
	title := fmt.Errorf("server %s failed to respond after %s", url, timeout)
	fmt.Println(title)
}
