// Упражнение 8.7 - программа параллельно скачивает все найденные страницы
// указанного в аргументе сайта. Не выполнена часть, в которой изменяются
// значения ссылочных html узлов
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

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int

	// запуск с аргументом коммандной строки
	n++
	go func() { worklist <- os.Args[1:] }()

	//!+ Параллельное сканирование и вызов функции сохранения
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
	//!-
}

func crawl(url string) []string {
	fmt.Println(url)     // весь вывод программы в этой функции - выводит полученный url
	tokens <- struct{}{} // захват маркера
	if err := save(url); err != nil {
		log.Fatal(err)
	}

	list, err := links.Extract(url) // получает со страницы список ссылок в абсолютном формате
	<-tokens                        // освобождение маркера
	if err != nil {
		log.Print(err)
	}

	return list // возвращает полученный список
}

// в этой переменной будет сохраняться домен, дальше которого качать html нельзя
var site string

// rawurl - исходный url
// пример: http://golang.org/doc/coryright.html
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
