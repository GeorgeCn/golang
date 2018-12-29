package main

import (
	"net"
	"io"
	"log"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy()
		log.Println("done")
		done<- struct{}{} //指示主 goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done // 等待后台goroutine完成
}
