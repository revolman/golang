// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

func main() {
	breadthFirst(crawl, os.Args[1:]) // в аргумента указывается ссылка на корневую страницу
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // отображение для проверки ссылок на повторяемость
	for len(worklist) > 0 {       // пока полчаемый список ссылок не равен нулю
		items := worklist            // переписывание полученного списка из worklist в items
		worklist = nil               // очистка worklist для наполнения его данными с новой страницы
		for _, item := range items { // очерёдный проход по всем полученным на предыдущей итерации ссылкам
			if !seen[item] { // если ссылка уже использовалась, то пропустить её
				seen[item] = true                       // в другом случае отметить её как использованную
				worklist = append(worklist, f(item)...) // f(item) = crawl(item) просканировать эту страницу
				// а список результатов добавить в worklist
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url) // весь вывод программы в этой функции - выводит полученный url
	if err := save(url); err != nil {
		log.Fatal(err)
	}

	list, err := links.Extract(url) // получает со страницы список ссылок в абсолютном формате
	if err != nil {
		log.Print(err)
	}

	return list // возвращает полученный список
}

// в этой переменной будет сохраняться домен, дальше которого качать html нельзя
var site string

// rawurl - исходный url
// http://golang.org/doc/coryright.html

func save(rawurl string) error {

	url, err := url.Parse(rawurl) // преобразование полученной строки в структуру *URL
	if err != nil {
		return fmt.Errorf("парсинг url: %v", err)
	}

	// назначить сайт. Если это первая итерация, то он будет равен аргументу
	if site == "" {
		site = url.Host
	}

	// если сайт уже был назначен и он отличен от нужного
	if site != url.Host {
		return nil
	}

	var dir string
	dir = url.Host // golang.org

	var filename string

	// Определяю путь и имя файла в разных случаях
	if strings.HasSuffix(url.Path, ".html") {
		parts := strings.Split(url.Path, "/")
		filename = parts[len(parts)-1]                      // copyright.html
		dir = dir + strings.Join(parts[:len(parts)-1], "/") // /doc
	} else {
		filename = "index.html"
		dir = dir + url.Path
	}

	fpath := filepath.Join(dir, filename)
	fmt.Println(fpath)

	// если директория не существует, то её нужно создать
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	// создать файл по определённому ранее пути
	file, err := os.OpenFile(fpath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("создание файла: %v", err)
	}
	defer file.Close()

	// делаю запрос к полученному url
	resp, err := http.Get(rawurl)
	if err != nil {
		return fmt.Errorf("запрос Get: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("получени тела овтета: %v", err)
	}

	if _, err := file.Write(body); err != nil {
		return fmt.Errorf("запись файла: %v", err)
	}

	return nil
}
