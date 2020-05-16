// Comma - разбивает числовую строку на разряды по 3 запятой
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("4111222333"))
}

func comma(s string) string {
	var buf bytes.Buffer
	for i := 0; i <= len(s)-1; i++ {
		if i == 0 {
			buf.WriteString(string(s[i]))
			continue
		}
		if len(s[i:])%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(s[i]))
	}
	return buf.String()
}
