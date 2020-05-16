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
		fmt.Println("Необходимо указвать адрес сайта и теги.")
		os.Exit(1)
	}

	url := os.Args[1]
	tags := os.Args[2:]

	outline(url, tags)
}

func outline(url string, tags []string) error {
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
	result := ElementByTagName(doc, tags...)
	images := ElementByTagName(doc, "img")
	heading := ElementByTagName(doc, "h1", "h2", "h3", "h4")

	for _, image := range images {
		fmt.Println("Найдено: ", image.Data)
	}

	for _, header := range heading {
		fmt.Println("Найдено: ", header.Data)
	}

	//!-call
	if result == nil {
		fmt.Printf("Элементы %q не найдены в HTML-файле сайта %s.\n", tags, url)
		os.Exit(1)
	}
	for _, item := range result {
		fmt.Printf("Найдено соответствие для %s в узле: %s\n", tags, item.Data)
	}
	return nil
}

// var depth int

// ElementByTagName ...
func ElementByTagName(doc *html.Node, tags ...string) (result []*html.Node) {
	// замена функции startElement
	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, tag := range tags {
				if n.Data == tag {
					result = append(result, n)
				}
			}
			// fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			// depth++
		}
	}
	// замена функции endElement
	// post := func(n *html.Node) {
	// 	if n.Type == html.ElementNode {
	// 		for _, tag := range tags {
	// 			if n.Data == tag {
	// 				result = append(result, n)
	// 			}
	// 		}
	// 		// depth--
	// 		// fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	// 	}
	// }

	forEachNode(doc, pre, nil)
	return
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, nil)
	}

	if post != nil {
		post(n)

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
