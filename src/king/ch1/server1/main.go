package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)// 回声请求调用处理程序
	log.Fatal(http.ListenAndServe("localhost:8500", nil))
}

//处理程序回显请求URL r的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
