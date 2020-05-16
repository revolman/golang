// Упражнение 1.4. Моя костыльная модификация dup2
// Печатает имена файлов, в которых найдены совпадения
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("Совпадений: %d, строки: %s\n", n, line)
			}
		}
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("[Dup2]: Ошибка: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			for line, n := range counts {
				if n > 1 {
					fmt.Printf("Совпадений: %d, строки: %s, файл: %s\n", n, line, arg)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
