// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.

// Алгоритм я подсмотрел у renatofq/gopl
package main

import (
	"fmt"
	"os"
	"sort"
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

	"data structures":       {"discrete math", "compilers"}, // "compilers" добавлины для создания цикла
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
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string                 // результирующий срез
	var cycler = make(map[string]bool) // проверка цикличности, по аналогии с seen, но с очисткой по возвращению
	seen := make(map[string]bool)      // проверка повторяемости

	var visitAll func(items []string) error // функция сортировки
	visitAll = func(items []string) error {
		for _, item := range items { // получаю ключи и для каждого ключа делаю действие
			// если cycler истинный, то это значит, что item уже был на этом этапе рекурсии
			if cycler[item] {
				return fmt.Errorf("%s", item)
			}
			// счётчики: были ли уже замечен текущий элемент
			if !seen[item] {
				cycler[item] = true
				seen[item] = true

				// проваливаемся в рекурсию и обрабатываем ошибку, если cycler будет true
				if err := visitAll(m[item]); err != nil {
					return fmt.Errorf("%s --> %s", item, err)
				}

				// очистка значения по возвращению из рекурсии обязательна!
				// иначе программа будет считать, что есть циклы там где их нет.
				cycler[item] = false

				// дополнить список
				order = append(order, item)
			}
		}
		return nil // ошибки нет
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, fmt.Errorf("Обнаружен цикл записимостей: %s", err)
	}

	return order, nil
}

//!-main
