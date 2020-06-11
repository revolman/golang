// Упражнение 8.6 - Ограничение глубины сбора ссылок на веб-странице.
package main

import (
	"flag"
	"fmt"
	"golang/ch5/links"
	"log"
	"os"
	"strings"
)

// Абстрактное представление ссылки и её "глубины".
type crawlList struct {
	list  []string
	depth int
}

// Буфферизированный канал, позволяет запускать не более 20 goroutine одновременно.
// Для борьбы с излишней параллельностью.
var tokens = make(chan struct{}, 20)

func crawl(url string, depth int) *crawlList {
	var cl crawlList
	fmt.Println(depth, url)
	tokens <- struct{}{} // Захват маркера
	list, err := links.Extract(url)
	<-tokens // Освобождение маркера
	if err != nil {
		log.Print(err)
	}
	cl.list = list
	cl.depth = depth + 1

	return &cl
}

func main() {
	var maxDepth int
	var list []string
	flag.IntVar(&maxDepth, "depth", 3, "максимальная глубина сканирования")
	flag.Parse()

	if strings.Contains(os.Args[1], "depth") {
		list = os.Args[2:]
	} else {
		list = os.Args[1:]
	}

	worklist := make(chan *crawlList)
	var n int // Количество ожидающих отправки в список
	// Запуск с аргументами командной строки.
	n++
	go func() {
		root := &crawlList{
			list:  list,
			depth: 1,
		}
		worklist <- root
	}()

	//  Паралелльное сканирование веб.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		cl := <-worklist
		if cl.depth > maxDepth {
			continue // новые элементы не попадут в crawl
		}
		for _, link := range cl.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link, cl.depth)
				}(link)
			}
		}
	}
}
