// Comma - разбивает числовую строку на разряды по 3 запятой
package main

import (
	"fmt"
)

func main() {
	fmt.Println(comma("12356789")) // "main"
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] // рекурсивный вызов функции с подстрокой за исключением последних трёх символов
}
