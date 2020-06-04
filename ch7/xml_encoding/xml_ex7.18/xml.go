package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Node ...
type Node interface{} // CharData or *Element

// CharData значение элемента CharData
type CharData string

// Element тип элемента с аттрибутами и указанием потомков
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

// переменная для корневого элемента

func main() {
	// !+debug
	file, err := os.Open("/home/revolman/git/golang/src/golang/ch7/xml_encoding/xml_ex7.18/index")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	// !-debug

	var genTree func(*Element) // определяется, что бы можно было использовать значение-функцию рекурсивно
	dec := xml.NewDecoder(file)
	var rootElement = new(Element)

	genTree = func(parent *Element) {
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
				el := &Element{ // описать полученный элемент как тип *Element
					Type: tok.Name,
					Attr: tok.Attr,
				}
				parent.Children = append(parent.Children, el) // Записать текущий элемент в срез потомков родительсткого элемента
				genTree(el)                                   // запустить рекурсивную обработку для текущего элемента, теперь он родитель
			// если закрывающий элемент, то вернуться на предыдущий уровень рекурсии
			case xml.EndElement:
				return

			// если получен текстовый элемент, то его значение подогнать в соответсвии с интерфейсом Node - преобразовать в тип CharData
			case xml.CharData:
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}

	// построить деревоP
	genTree(rootElement)
	// вывод дерева
	PrintTree(rootElement, 0)
	// printElement(rootElement)

}

// func printElement(e *Element) {
// 	var printElementRec func(*Element, string)

// 	printElementRec = func(e *Element, prefix string) {
// 		fmt.Printf("%s%s\n", prefix, e.Type.Local)
// 		prefix += "  "
// 		for _, node := range e.Children {
// 			switch node := node.(type) {
// 			case *Element:
// 				printElementRec(node, prefix)
// 			case CharData:
// 				fmt.Printf("%s%s", prefix, node)
// 			}
// 		}
// 	}

// 	printElementRec(e, "")
// }

// PrintTree отображает дерево элементов с отступами
func PrintTree(n Node, depth int) {
	switch n := n.(type) {
	case *Element:
		if n.Type.Local != "" {
			fmt.Printf("%*s%s\n", depth*2, "", n.Type.Local)
		}
		for _, children := range n.Children {
			depth++
			PrintTree(children, depth)
		}
	case CharData:
		fmt.Printf("%*s%s\n", depth*2, "", n)

	default:
		panic(fmt.Sprintf("Получен неожиданный тип %T", n))
	}
}

// GenTree генерирует дерево элементов xml файла
// func GenTree(parent *Element, file *os.File) {

// 	// defer file.Close()
// 	// dec := xml.NewDecoder(os.Stdin)
// 	dec := xml.NewDecoder(file)

// 	for {
// 		tok, err := dec.Token()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
// 			os.Exit(1)
// 		}

// 		switch tok := tok.(type) {

// 		case xml.StartElement:
// 			el := &Element{ // описать полученный элемент как тип *Element
// 				Type: tok.Name,
// 				Attr: tok.Attr,
// 			}
// 			parent.Children = append(parent.Children, el) // Записать текущий элемент в срез потомков родительсткого элемента
// 			GenTree(el, file)                             // запустить рекурсивную обработку для текущего элемента, теперь он родитель
// 		// если закрывающий элемент, то вернуться на предыдущий уровень рекурсии
// 		case xml.EndElement:
// 			return

// 		// если получен текстовый элемент, то его значение подогнать в соответсвии с интерфейсом Node - преобразовать в тип CharData
// 		case xml.CharData:
// 			parent.Children = append(parent.Children, CharData(tok))
// 		}
// 	}
// }
