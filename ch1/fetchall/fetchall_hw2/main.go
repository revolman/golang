// fetchall выполняет параллельную выборку URL и сообщает
// о затраченном времени и развере ответа для каждого из них
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.OpenFile("fetchall.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Упражнение 1.10 - запись результатов в файл
	if err != nil {                                                                     // Обработка ошибок
		fmt.Println("Не получилось создать файл: \n", err)
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http") != true { // Проверка префикса http*
			url = "https://" + url // Добаление префикса https://
		}
		go fetch(url, ch) // запуск go-подпрограммы
	}
	for range os.Args[1:] {
		file.WriteString(<-ch)
		// Получение из канала ch
	}
	file.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s \n", err) // Отправка в канал ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d, %s\n", secs, nbytes, url)
}
