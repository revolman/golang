// fetchall выполняет параллельную выборку URL и сообщает
// о затраченном времени и развере ответа для каждого из них
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // запуск go-подпрограммы
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		// Получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка в канал ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // (dst псевдо-writer, src Reader)
	// возвращает количество записанных байт и error, т.е. само тело ответа отбрасывается.
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d, %s", secs, nbytes, url)
}
