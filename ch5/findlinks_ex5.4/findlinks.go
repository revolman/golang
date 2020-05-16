// Findlinks1 выводит ссылки в HTML-документе,
// прочитанном со стандартного входа.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
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

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		links = linker(links, n, "a", "href")
		links = linker(links, n, "link", "href")
		links = linker(links, n, "script", "src")
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func linker(links []string, n *html.Node, dataType string, linkType string) []string {
	if n.Data == dataType {
		for _, a := range n.Attr {
			if a.Key == linkType {
				links = append(links, a.Val)
			}
		}
	}
	return links
}
