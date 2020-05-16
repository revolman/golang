// Напишите функцию expand(s string, f func(string) string) string,
// которая заменяет каждую подстроку $foo в s текстом, который
// возвращается вызовом f("foo")

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "Это тест $foo $foo $foo \n"

	// не понял сути задачи, f не нужна.
	// предположим, что f должна обрабатывать текст - выводить прописными.
	f := func(s string) string {
		return strings.ToUpper(s)
	}
	fmt.Print(expand(s, f))
}

func expand(s string, f func(string) string) string {
	if len(os.Args[1:]) < 2 {
		fmt.Printf("Аргументы: [\"стараная строка\"] [новая строка]\n" +
			"Если старая строка состоит из одного слова, то кавычки можно опустить")
	}

	oldString := os.Args[1]
	newString := strings.Join(os.Args[2:], " ")

	replaced := strings.Replace(s, oldString, f(newString), -1)

	return replaced
}
