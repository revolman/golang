// Упражнение 5.17 - вариативная версия strings.Join
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("Нужно указать разделитель и текст")
		os.Exit(1)
	}

	sep := os.Args[1]
	str := os.Args[2:]
	fmt.Println(concat(sep, str...))
}

func concat(sep string, elements ...string) string {
	//!+Поигрался с логами
	log.SetPrefix("concat: ")
	log.SetFlags(0)
	//!-Поигрался с логами
	if len(elements) == 0 {
		log.Fatal("не указаны аргументы")
	}

	var result string
	for _, element := range elements {
		result = result + sep + element
	}
	return result
}
