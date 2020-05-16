package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // множество строк
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true // в отображение записывается ключ - строка и её значение меняется на true
			fmt.Println(line) // и только в этом случае будет сделан вывод этой строки
		}
	}
	if err := input.Err(); err != nil { // обработка ошибок ввода
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
