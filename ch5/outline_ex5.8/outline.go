// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args[1:]) < 2 {
		fmt.Println("Необходимо указвать адрес сайта и тег.")
		os.Exit(1)
	}

	url := os.Args[1]
	id := os.Args[2]

	outline(url, id)
}

func outline(url string, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	resultNode := ElementByID(doc, id)
	//!-call
	if resultNode == nil {
		fmt.Printf("Элемент с id %q не найден в HTML-файле сайта %s.\n", id, url)
		os.Exit(1)
	}
	fmt.Printf("Найден первый элемент %q в узле:\n%#v\n", id, resultNode)

	return nil
}

// var depth int

// ElementByID ...
func ElementByID(doc *html.Node, id string) (result *html.Node) {
	// замена функции startElement
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					result = n
					return true
				}
			}
			// fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			// depth++
		}
		return false
	}
	// замена функции endElement
	post := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					result = n
					return true
				}
			}
			// depth--
			// fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
		return false
	}

	forEachNode(doc, pre, post)
	return
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if post(n) {
			return
		}
	}
}

// var depth int

// func startElement(n *html.Node, id string) bool {
// 	if n.Data == id { return true }
// 	if n.Type == html.ElementNode {
// 		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
// 		depth++
// 	}
// 	return false
// }

// func endElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		depth--
// 		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
// 	}
// }
