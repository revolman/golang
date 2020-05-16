package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // отображение для проверки ссылок на повторяемость
	for len(worklist) > 0 {       // пока полчаемый список ссылок не равен нулю
		items := worklist            // переписывание полученного списка из worklist в items
		worklist = nil               // очистка worklist для наполнения его данными с новой страницы
		for _, item := range items { // очерёдный проход по всем полученным на предыдущей итерации ссылкам
			if !seen[item] { // если ссылка уже использовалась, то пропустить её
				seen[item] = true // в другом случае отметить её как использованную

				// f(item)... - см. стр. 177. т.к. append() является вариативной функцией,
				// то аргументы, которые уже находятся в стрезе, а в данному случае функция возвращает срез []string,
				// нужно передавать указав ... после последнего элемента.
				// append(worklist, f(item)...) равноценно append(worklist, "path1", "path2", "pathN"})
				worklist = append(worklist, f(item)...) // f(item) = crawl(item) просканировать эту страницу
				// а список результатов добавить в worklist
			}
		}
	}
}

func crawl(path string) []string {
	fmt.Println(path) // вывод пути

	list, err := listDir(path)
	if err != nil {
		log.Print(err)
	}

	return list // возвращает полученный список
}

// listDir - получает список директорий
// возвращает список вложенных директорий в абсолютном формате
func listDir(path string) ([]string, error) {
	var result []string

	fileinfo, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("получение fileinfo: %s", err)
	}

	if !fileinfo.IsDir() {
		return nil, fmt.Errorf("не директория")
	}

	list, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("чтение директории: %s", err)
	}

	for _, item := range list {
		if !item.IsDir() {
			continue
		}
		fpath := filepath.Join(path, item.Name())
		result = append(result, fpath)
	}
	return result, nil
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
