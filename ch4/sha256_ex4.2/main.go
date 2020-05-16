// Упражнение 4.2 - по умолчанию выводит sha256, но при заданных флагах используются другие алгаритмы
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Определяю флаги
var sha3 = flag.String("sha384", "", "-sha384 \"long arg\"")
var sha5 = flag.String("sha512", "", "-sha512 \"long arg\"")

func main() {
	var input string

	arg := os.Args[1:] // сохраняю введённые аргументы

	if len(arg) < 1 {
		fmt.Println("Необходимо ввести строку")
		os.Exit(1)
	}

	flag.Parse() // считываю флаги

	if len(*sha3) > 0 { // при наличии флага -sha384
		input = *sha3
		fmt.Printf("%s\n=\n%x\n", input, sha512.Sum384([]byte(input)))
		os.Exit(0)
	}

	if len(*sha5) > 0 { // при наличии флага -sha512
		input = *sha5
		fmt.Printf("%s\n=\n%x\n", input, sha512.Sum512([]byte(input)))
		os.Exit(0)
	}

	input = fmt.Sprint(strings.Join(arg, " ")) // разделяю аргументы пробелами, форматирую их как текст

	fmt.Printf("%s\n=\n%x\n", input, sha256.Sum256([]byte(input))) // получаю дайджест
}
