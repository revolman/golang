package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// Element представляет имя элемента и его аттрибуты
type Element struct {
	name string
	attr []Attribute
}

// Attribute представляет аттрибут типа ключ - значение
type Attribute struct {
	name  string
	value string
}

// Пример:
// <div id="one">
// 	<div id="two">
// 		<p>Hello</p>
// 	</div>
// 	<h1>
// </div>

// Match позволяет сравнивать тип Attribute
func (a Attribute) Match(j Attribute) bool {
	if a.name == j.name && a.value == j.value {
		return true
	}
	return false
}

// String - выводит только список аргументов
func (e Element) String() string {
	var str string
	for i := 0; i < len(e.attr); i++ {
		str += fmt.Sprintf("%s=%q", e.attr[i].name, e.attr[i].value)
	}
	return str
}

// Match позволяет сравнивать тип Element
func (e Element) Match(j Element) bool {
	if e.name == j.name && len(e.attr) != 0 { // добавил проверку количества аргументов
		for i := 0; i < len(j.attr); i++ {
			if !e.attr[i].Match(j.attr[0]) {
				return false
			}
			return true
		}
	} else if e.name == j.name {
		return true // если имена совпали, а аргументов нет
	}
	return false
}

func main() {
	input := inputParse()

	// Тут инфа для дебага
	// for _, e := range input {
	// 	fmt.Printf("Список элементов: %s\nАттрибуты: ", e.name)

	// 	for _, attr := range e.attr {
	// 		fmt.Printf("%s=%q ", attr.name, attr.value)
	// 	}
	// 	fmt.Println("")
	// }
	// }

	dec := xml.NewDecoder(os.Stdin)

	var stack []Element

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
			var currentItem Element

			// запись аттрибутов элемента
			for _, attr := range tok.Attr {
				var attribute Attribute
				attribute.name = attr.Name.Local
				attribute.value = attr.Value
				currentItem.attr = append(currentItem.attr, attribute)
			}
			// запись имени элемента
			currentItem.name = tok.Name.Local
			// добавление элемента в стек
			stack = append(stack, currentItem)

		case xml.EndElement:
			stack = stack[:len(stack)-1] // снятие со стека

		case xml.CharData:
			// сравнить содержатся ли элементы intput в stack в том же порядке
			if containsAll(stack, input) {
				var s []string
				for _, item := range stack {
					s = append(s, item.name)
					// fmt.Printf("%s%s", item.name, item.String())
				}
				fmt.Printf("%s: %s\n", strings.Join(s, " "), tok)
				// fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok.Name)
			}
		}
	}
}

// xml.go "div id=one" "div id=two"
// подготовка входных данных
func inputParse() []Element {
	var result []Element
	input := os.Args[1:]

	for _, s := range input {
		// объявление нового элемента и его аттрибута
		var input Element
		var attribute Attribute

		args := strings.Split(s, " ")

		for _, arg := range args {
			if strings.Contains(arg, "=") {
				q := strings.Split(arg, "=")
				attribute.name = q[0]
				attribute.value = q[1]
				input.attr = append(input.attr, attribute)
				continue
			}
			input.name = arg
			// нужна проверка на случай, если будет введён некорректный тег
			// например xml.go "div div id=one". Разрешить указывать только одно имя тега.
		}
		result = append(result, input)
	}
	return result
}

// Проверка, содежатся ли элементы input в stack в том же порядке.
func containsAll(stack, input []Element) bool {

	for len(input) <= len(stack) {
		if len(input) == 0 {
			return true
		}
		if stack[0].Match(input[0]) {
			input = input[1:]
		}
		stack = stack[1:]
	}
	return false
}

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
// }
