// Comma - Проверяет являются ли входящие строки анаграммами.
// Упражнение 3.12
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalln("Usage: comma [arg1 string] [arg2 string]")
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	if anagramma(arg1, arg2) {
		fmt.Println("Это анаграмма!")
	} else {
		fmt.Println("Это не анаграмма!")
	}
}

// anagramma - проверяет являются ли строки анограммами друг друга.
// Для проверки подсчитываю сумму кодов UTF-8 каждой руны в строке, затем сравниваю.
func anagramma(s1 string, s2 string) bool {
	var sum1 int
	var sum2 int
	if s1 == s2 { // Вообще, если строки одинаковы, то я бы сказал, что это частный случай анаграммы, но ОК.
		return false
	}
	for _, r := range s1 {
		sum1 += int(r)
	}
	for _, r := range s2 {
		sum2 += int(r)
	}
	if sum1 == sum2 {
		return true
	}
	return false
}
