package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type element struct {
	name string
	attr []attrib
}

type attrib struct {
	name  string
	value string
}

// xml.go div id="page" class="wide"
// подготовка входных данных
func inputParse() []element {
	var input []element

	var attribute attrib
	in := os.Args[1:]

	for _, arg := range in {
		var el element

		if strings.Contains(arg, "=") {
			q := strings.Split(arg, "=")
			attribute.name = q[0]
			attribute.value = q[1]
			el.attr = append(el.attr, attribute)
			continue
		}
		el.name = arg
		input = append(input, el)
	}
	return input
}

func main() {
	in := inputParse()
	for _, el := range in {
		fmt.Printf("Элемент: %s\nАттрибуты: ", el.name)
		for _, attr := range el.attr {
			fmt.Printf("%s=%q ", attr.name, attr.value)
		}
	}
	fmt.Println("")
	// }

	dec := xml.NewDecoder(os.Stdin)

	var stack []element

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			var processingElement element

			for _, attr := range tok.Attr {
				var attributes attrib
				attributes.name = attr.Name.Local
				attributes.value = attr.Value
				processingElement.attr = append(processingElement.attr, attributes)
			}

			processingElement.name = tok.Name.Local

			stack = append(stack, element{name: tok.Name.Local}) // запись в стек

		case xml.EndElement:
			stack = stack[:len(stack)-1] // снятие со стека

		case xml.CharData:
			// принимать аттрибуты в виде div id="page"
			if containsAll(stack, in) {
				// fmt.Printf("%s with attr: %s=%q: : %s\n", strings.Join(stack, " "), tok.Name)
			}
		}
	}
}

// // type StartElement struct {
// // 	Name Name
// // 	Attr []Attr
// // }

// // type Attr struct {
// // 	Name  Name
// // 	Value string
// // }

// // type Name struct {
// // 	Space, Local string
// // }

func containsAll(stack []element, input []element) bool {
	for _, el := range stack {

	}

	return false

	// for len(y) <= len(x) {
	// 	if len(y) == 0 { // когда в y ни чего не останется вернуть истину
	// 		return true
	// 	}
	// 	if x[0] == y[0] {
	// 		y = y[1:] // если совпали, то перейти к следующему y
	// 	}
	// 	x = x[1:] // перейти к следующему x
	// }
	// return false
}

func compareItems(i, j element) bool {
	if i.name == j.name {
		if len(j.attr) == 0 {
			return true
		}
		if len(j.attr) <= len(i.attr) {

		}
	}

	return false
}
