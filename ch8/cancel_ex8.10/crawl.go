// Упражнение 8.10 - Отмена HTTP запросов посредством закрытия
// необязательного канала Cancel в структуре http.Request.
package main

import (
	"fmt"
	"log"
	"os"
)

// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества паралленльных запросов величиной 20.
var tokens = make(chan struct{}, 20)

// cancel это канал отмены запроса
var cancel = make(chan struct{})

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // Захват маркера
	list, err := Extract(url, cancel)
	<-tokens // Освобождение маркера
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // Количество ожидающих отправки в список

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("Interrupted")
		close(cancel)
	}()

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
