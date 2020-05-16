package main

import (
	"fmt"
)

func main() {
	var runes []rune
	for _, r := range "Hello, world!" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
