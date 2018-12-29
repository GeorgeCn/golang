package main

//Findlinks1 输出从标准输入中读入的HTML文档中的所有链接
import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//visit 函数会将 n 节点中的每个链接添加到结果中
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c:= n.FirstChild; c !=nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
