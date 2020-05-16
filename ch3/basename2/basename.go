// Basename - аналог утилиты basename
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("././././asd.fg.123")) // "asd.fg"
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	// dot := strings.LastIndex(s, ".")  // можно и так
	// s = s[:dot]
	if dot := strings.LastIndex(s, "."); dot >= 0 { // пример
		s = s[:dot]
	}
	return s
}
