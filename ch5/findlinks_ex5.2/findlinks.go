// Упражнение 5.2 - напишите функцию для заполнения отображения, ключами которого
// являются имена элементов (p, div, span и т.д.), а значениями - количество элементов
// с таким именем в дереве HTML-документа.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var count = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)

	for k, v := range count {
		fmt.Printf("Тег %s встречается %d раз\n", k, v)
	}
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	// далее рекурсивная обработка потомков и соседей
	if n.FirstChild != nil {
		visit(n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling)
	}
}
