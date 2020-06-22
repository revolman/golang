// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества паралленльных запросов величиной 20.
package main

import (
	"fmt"
	"github/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // Захват маркера
	list, err := links.Extract(url)
	<-tokens // Освобождение маркера
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // Количество ожидающих отправки в список

	// Запуск с аргументами коммандной строки.
	n++
	go func() { worklist <- os.Args[1:] }()


	//  Паралелльное сканирование веб.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
