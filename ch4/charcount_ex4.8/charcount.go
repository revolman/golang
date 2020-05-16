// Упражнение 4.8 - версия charcount, которая дополнительно подсчитывает
// категории unicode символов, а именно количество букв и цифр в тексте
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	var invalid, letter, digit int

	in := bufio.NewReader(os.Stdin)
	for { // бесконечный цикл
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
		}
		if unicode.IsLetter(r) {
			letter++
		}
		if unicode.IsDigit(r) {
			digit++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Println("\nlen\tcount")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if letter > 0 {
		fmt.Printf("\n%d символов букв в тексте.\n", letter)
	}
	if digit > 0 {
		fmt.Printf("%d цифр в тексте.\n", digit)
	}
	if invalid > 0 {
		fmt.Printf("%d неверных символов UTF-8\n", invalid)
	}
}
