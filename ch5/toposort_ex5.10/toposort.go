// Упражнение 5.10 - вместо срезов использовать отображения
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for course, dependencies := range topoSort(prereqs) {
		fmt.Printf("%s:\t%s\n", course, dependencies)
	}
}

func topoSort(m map[string][]string) map[string][]string {
	// результирующий срез
	// var order []string
	var result = make(map[string][]string)
	// проверка на повторяемость
	seen := make(map[string]bool)

	var visitAll func(items []string) []string // функция будет возвращать список зависимостей для каждого ключа
	visitAll = func(items []string) []string {
		var depends []string
		for _, item := range items { // для каждого значения переданного ключа
			if !seen[item] { // проверить повторяемость. Держать в голве, что этот список тоже нужно обнулить для каждого ключа!!!
				seen[item] = true
				depends = append(depends, item)                 // добавить само значение, т.к. это первая зависимость
				depends = append(depends, visitAll(m[item])...) // добавить зависимости зависимости
			}

		}
		return depends
	}

	// список ключей
	// var keys []string

	// применение функции к каждому ключу
	for k, value := range m {
		result[k] = visitAll(value)
		seen = make(map[string]bool) // обнуление счётчика проверки повторяемости для каждого ключа
	}
	// предварительная сортировка
	// sort.Strings(keys)
	// вызов функции
	// visitAll(keys)

	// Если значение ключа в своём значении имеет ключ, тогда сообщить о цикле

	// результат
	return result
}

//!-main
