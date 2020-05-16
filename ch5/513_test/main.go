package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "https://golang.org/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// parts := strings.Split(url, "/")
	// fpath := strings.Join(parts[2:], "/")
	fpath := strings.TrimLeft(url, "https://")
	fname := "page.html"
	fullName := fpath + fname

	fmt.Println(fpath + fname)
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		os.MkdirAll(fpath, 0755)
	}

	file, err := os.OpenFile(fullName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("создание файла: %v", err)
	}
	defer file.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("получени тела овтета: %v", err)
	}

	if _, err := file.Write(body); err != nil {
		log.Fatalf("запись файла: %v", err)
	}
}
