// Упражнение 5.3 - напишите функцию для вывода содержимого всех текстовых узлов
// в дереве документа HTML. Не входите в элементы <script> и <style>.
// более-менее сносное решение
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	buf := &bytes.Buffer{}
	collectText(doc, buf)
	fmt.Println(buf)

	// reader := bufio.NewReader(buf)
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	trimed := strings.TrimSpace(line)
	// 	if len(trimed) != 0 {
	// 		fmt.Println(trimed)
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			trimmed := strings.TrimSpace(n.Data)
			if len(trimmed) != 0 {
				buf.WriteString(n.Data + "\n")
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}
