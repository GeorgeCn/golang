package main

import (
	"os"
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

func main() {
	for _, url := os.Args[1:]
	links, err := findLinks(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		continue
	}
	for _, link := range links {
		fmt.Println(link)
	}
}

// findLinks 发起一个Http的GET请求， 解析返回的HTML页面，并返回所有链接
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}
