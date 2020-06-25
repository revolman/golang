// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества паралленльных запросов величиной 20.
package main

import (
	"fmt"
	"golang/ch5/links"
	"log"
	"os"
	"sync"
)

var tokens = make(chan struct{}, 20)

func cancelled() bool {
	select {
	case <-Cancel:
		return true
	default:
		return false
	}
}

func crawl(url string, wg *sync.WaitGroup) []string {
	defer wg.Done()
	if cancelled() {
		return nil
	}
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
	var wg sync.WaitGroup

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(Cancel)
	}()

	go func() {
		wg.Wait()
		close(worklist)
	}()

	// Запуск с аргументами коммандной строки.
	n++
	go func() {
		wg.Add(1)
		worklist <- os.Args[1:]
	}()

	//  Паралелльное сканирование веб.
	seen := make(map[string]bool)

loop:
	for {
		select {
		case <-Cancel:
			for range worklist {
				// ни чего не делать
			}
			return
		case list, ok := <-worklist:
			if !ok {
				break loop
			}
			for ; n > 0; n-- {
				// list := <-worklist
				for _, link := range list {
					if !seen[link] {
						seen[link] = true
						n++
						go func(link string) {
							wg.Add(1)
							worklist <- crawl(link, &wg)
						}(link)
					}
				}
			}
		}
	}
	fmt.Println("Interrupted!")
}
