// Comma - разбивает числовую строку на разряды по 3 запятой
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-12235.6789998"))
}

func beforeComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return beforeComma(s[:n-3]) + "," + s[n-3:] // рекурсивный вызов функции с подстрокой за исключением последних трёх символов
}

func afterComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:3] + "," + afterComma(s[3:]) // рекурсивный вызов функции с подстрокой за исключением последних трёх символов
}

func comma(s string) string {
	var buf bytes.Buffer
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
	}
	dot := strings.Index(s, ".") // вернёт -1, если не находит строку
	if dot == -1 {
		buf.WriteString(beforeComma(s))
	} else {
		buf.WriteString(beforeComma(s[1:dot]))
		buf.WriteByte('.')
		buf.WriteString(afterComma(s[dot+1:]))
	}
	return buf.String()
}
