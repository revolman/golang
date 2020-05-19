package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// Reader - my reader struct
type Reader struct {
	s string
	// i        int64 // current reading index
	// prevRune int   // index of previous rune; or < 0
}

// Read - метод для соответствия контракту io.Reader
func (r *Reader) Read(b []byte) (n int, err error) {
	n = copy(b, r.s) // Записываем байты в структуру, возвращается количество
	r.s = r.s[n:]    // то что прочитали нам уже не нужно, перезаписываем. Можно обойтись без этого
	if len(r.s) == 0 {
		err = io.EOF // EOF если в не осталось символов
	}
	return
}

// NewReader ...
// что бы вернуть io.Reader нужно создать метод Read для типа Reader
func NewReader(s string) io.Reader {
	return &Reader{s}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode { // кастрировал программу для проверки
		// for _, a := range n.Attr {
		// 	if a.Key == "href" {
		links = append(links, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	buf := &bytes.Buffer{}
	reader := NewReader("Test\n")
	buf.ReadFrom(reader)
	fmt.Printf("%s", buf.String())

	s := "<html><body><a href=><a/><div><span></span></div><body></html>"
	doc, err := html.Parse(NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "getting node: %s", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Printf("<%s>\n", link)
	}
}
