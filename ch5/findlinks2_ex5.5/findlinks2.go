package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var freq = make(map[string]int)

func main() {

	if len(os.Args[1:]) > 1 {
		fmt.Println("В качестве аргумента нужно указывать корректный URL")
		os.Exit(1)
	}
	url := os.Args[1]

	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal("Получена ошибка: ", err)
	}

	fmt.Printf("Слов: %d\nИзображений: %d\n", words, images)
}

// CountWordsAndImages ...
func CountWordsAndImages(url string) (words, images int, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Ошибка парсинга HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			freq["images"]++
		}
	}
	if n.Type == html.TextNode {
		freq["words"] += wordsCounter(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c)
	}

	words = freq["words"]
	images = freq["images"]

	return
}

func wordsCounter(s string) int {
	var counter int
	reader := strings.NewReader(s)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counter++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return counter
}
