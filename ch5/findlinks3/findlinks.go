// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // отображение для проверки ссылок на повторяемость
	for len(worklist) > 0 {       // пока полчаемый список ссылок не равен нулю
		items := worklist            // переписывание полученного списка из worklist в items
		worklist = nil               // очистка worklist для наполнения его данными с новой страницы
		for _, item := range items { // очерёдный проход по всем полученным на предыдущей итерации ссылкам
			if !seen[item] { // если ссылка уже использовалась, то пропустить её
				seen[item] = true                       // в другом случае отметить её как использованную
				worklist = append(worklist, f(item)...) // f(item) = crawl(item) просканировать эту страницу
				// а список результатов добавить в worklist
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)                // весь вывод программы в этой функции - выводит полученный url
	list, err := links.Extract(url) // получает со страницы список ссылок в абсолютном формате
	if err != nil {
		log.Print(err)
	}
	return list // возвращает полученный список
}

func main() {
	breadthFirst(crawl, os.Args[1:]) // в аргумента указывается ссылка на корневую страницу
}
