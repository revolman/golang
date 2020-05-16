// Упражнение 4.9 - подсчёт частоты каждого слова во входном текстовом файле.
// Либо может принимать обычный ввод stdin из консоли, для подтверждения CTRL+D
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		freq[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("word\tcount")
	for k, v := range freq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
