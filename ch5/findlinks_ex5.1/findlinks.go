// Упражнение 5.1 - изменить программу так, что бы она обходила
// связанный список n.FirstChild с помощью рекурсивных вызовов
// visit, а не цикла
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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// если у узла есть потомок - провалиться в рекурсию до того момента,
	// как потомки не закончатся. Когда заканчиваются потомки поток возвращается
	// к тому моменту где было ветвеление, если оно было
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	// если было ветвление, значит у узла был сосед NextSibling, тогда поток нужно
	// направить в соседний узел, в котором всё точно так же рекурсивно повторится
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	// когда заканчиваются потомки и не было соседей возвращается результат.
	return links
}
