// Упражнение 5.7 Outline - Разработайте startElement и endElement для обобщенного вывода HTML.
// Выводите узлы комментариев, текстовые узлы и атрибуты каждого элемента (<a href='...'>).
// Используйте сокращенный вывод наподобие <img/> вместо <img></img>, когда элемент не имеет дочерних узлов.
// Напишите тестовую программу. 
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
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
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// продготавливаю строку из всех атрибутов элемента
		var ats string

		for _, a := range n.Attr {
			ats += fmt.Sprintf(" %s='...'", a.Key)
		}
		// если потомков нет, то выводить сокращённый тег
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, ats)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, ats)
		}
		depth++
	}

	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
		depth++
	}

	if n.Type == html.TextNode {
		// создаю буффер, разбиваю строки
		buf := &bytes.Buffer{}
		buf.WriteString(n.Data + "\n")
		reader := bufio.NewReader(buf)
		// обрабатываю каждую строку по отдельности, что бы выводить их с корректным отсутпом
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break // чтение завершено
			}
			if err != nil {
				log.Fatalf("startElement: reader: %v", err)
			}
			// опускаю строки с пробелами
			trimmed := strings.TrimSpace(line)
			if len(trimmed) != 0 {
				fmt.Printf("%*s%s\n", depth*2, "", trimmed)
			}
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		// если есть потомки, то выводить закрывающий тег
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
