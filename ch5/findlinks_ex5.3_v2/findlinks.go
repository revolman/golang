// Упражнение 5.3 - напишите функцию для вывода содержимого всех текстовых узлов
// в дереве документа HTML. Не входите в элементы <script> и <style>.
// Сомнительное решение, т.к. часть информации теряется.
package main

import (
	"fmt"
	"os"
	"unicode"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	collectText(doc)

}

func collectText(n *html.Node) {
	if n.Type == html.TextNode && n.Parent.Data != "script" {
		if noSpaces([]rune(n.Data)) {
			fmt.Print(n.Data, " ")
		}

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c)
	}
}

func noSpaces(s []rune) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return true
		}
	}
	return false
}
