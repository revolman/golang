// Basename - аналог утилиты basename
package main

import "fmt"

func main() {
	fmt.Println(basename("/home/revolman/main.go")) // "main"
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' { // доходит до знака /, и выводит всё что перед ним
			s = s[i+1:] // нужно добавить 1 к индексу, что бы не захватить /
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i] // здесь наоборот выводит всё что до точки, т.к. в записи s[:i] i-не включительно
		}
	}
	return s
}
