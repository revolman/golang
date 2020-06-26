// Упражнение 8.11 - Версия Fetch, которая выполняет запрос
// к нескольким URL, когда получен первый ответ остальные
// запросы отменяются
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	if len(urls) < 1 {
		fmt.Println("Usage: first [https://url ...]")
		os.Exit(0)
	}
	responses := make(chan *http.Response, len(urls))
	cancel := make(chan struct{})
	for _, url := range urls {
		go func(url string) {
			req, err := http.NewRequest("GET", url, nil)

			req.Cancel = cancel
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			responses <- resp
		}(url)
	}
	firstResp := <-responses
	b, err := ioutil.ReadAll(firstResp.Body)
	close(cancel)
	firstResp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
