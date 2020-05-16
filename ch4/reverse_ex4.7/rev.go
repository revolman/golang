// Упражненеи 4.7 - измененная версия reverse,
// обращает последовательность среза []byte без выделения доп памяти.

package main

import "fmt"

func main() {
	text := "Hello world!"
	s := []byte(text)
	fmt.Printf("%s\n%q\n", text, reverse(s))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
