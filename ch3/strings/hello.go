package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 可乐"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("i: %d\trune: %c\t size: %d\t\tlen(s): %d\n", i, r, size, len(s))
		i += size
	}
}
