package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	url := os.Args[1]
	fetch(url)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	// Закрыть файл по завершению функции, но предпочесть ошибку io.Copy ошибке f.Close,
	// т.к. в некоторых ФС фиксирование ошибок записи может быть отложено
	// до тех пор, пока файл не будет закрыт.
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}
