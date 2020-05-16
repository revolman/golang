// Упражнение 4.6 - преобразует последовательности смежных пробелов
// в срезе []byte в кодировке UTF-8 в один пробел ASCII.
// Использую алгоритм из упражнения 4.5

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	text := "Finally,    the    breakfast!  I've " +
		"been      waiting this for  so    long!"
	s := []byte(text)
	fmt.Printf("%q\n", spaces(s))
}

func spaces(slice []byte) []byte {
	var last rune
	var correction int

	for i, b := range slice {
		r, _ := utf8.DecodeRune(slice[i:])
		fmt.Printf("%[1]c\n", r)
		if unicode.IsSpace(r) && unicode.IsSpace(last) {
			// fmt.Println("Пробел. ID: ", i) // dubug =)
			correction++
		}
		last = r
		slice[i-correction] = b
	}

	return slice[:len(slice)-correction]
}
