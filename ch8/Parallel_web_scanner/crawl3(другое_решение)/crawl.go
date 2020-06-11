// Альтернативное решение проблемы черезмерного параллелизма.
// Версия crawl без подсчитывающего семафора вызывается одной из 20
// долгоживущих go-подпрограмм сканирования, гарантируя тем самым,
// что одновременно активны не более 20 HTTP-запросов.
package main

import (
	"fmt"
	"golang/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // список URL, могут быть дубли.
	unseenLinks := make(chan string) // Удаление дублей.

	go func() { worklist <- os.Args[1:] }()

	// Создание 20 го-подпрограмм сканирования для выборки
	// всех непросмотренных ссылок.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Главная го= подпрограмма удаляет дубликаты из списка
	// и отправляет непросмотренные ссылки сканерам.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
