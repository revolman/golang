package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Comics - структура для хранения информации о публикации
type Comics struct {
	Day        string
	Month      string
	Year       string
	Num        int
	IMG        string
	Title      string
	Transcript string
}

const xkcdURL string = "https://xkcd.com"

func main() {
	var fname = "comics.json"
	_, err := os.Stat(fname)

	if os.IsNotExist(err) {
		fmt.Println("Архива не обнаружен. Подготовка архива:")

		template, _ := json.MarshalIndent([]Comics{}, "", " ")
		if err := ioutil.WriteFile(fname, template, 0644); err != nil {
			log.Fatal("Ошибка при создании шаблона файла: ", err)
		}
		getComs(1, lastOnSite(), fname)
	} else {
		getComs(lastIndex(fname)+1, lastOnSite(), fname)
		fmt.Println("Архив а актуальном состоянии.")
	}

	if len(os.Args[1:]) < 2 {
		exitWithUsage()
	}

	// Вызов поиска
	if os.Args[1] == "search" {
		args := os.Args[2:]
		fmt.Printf("Поиск %q в архиве\n", strings.Join(args, " "))
		search(fname, args)
		os.Exit(0)
	}
	exitWithUsage()
}

func exitWithUsage() {
	fmt.Printf("Использование:\nsearch QUERY\t - поиск по архиву\n")
	os.Exit(0)
}

// search - обеспечивает поиск по архиву
func search(fname string, args []string) { // фиговенький алгоритм, нужно позже переписать.
	archive := alreadyInArchive(fname)
	query := strings.Join(args, " ")

	var result1, result2, result3 []*Comics

	for _, item := range archive {
		if strings.Contains(item.Transcript, query) {
			result1 = append(result1, item)
		}

		if len(args) >= 3 {
			if strings.Contains(item.Transcript, args[0]) && strings.Contains(item.Transcript, args[1]) && strings.Contains(item.Transcript, args[2]) {
				result2 = append(result2, item)
			}
		}

		if len(args) >= 2 {
			if strings.Contains(item.Transcript, args[0]) && strings.Contains(item.Transcript, args[1]) {
				result3 = append(result3, item)
			}
		}
	}

	fmt.Println("Точные совпадения:")
	for _, item := range result1 {
		fmt.Printf("#%-5d %-25.25s %-55.55s...\n", item.Num, xkcdURL+"/"+strconv.Itoa(item.Num), item.Transcript)
	}
	fmt.Println("Совпадение по трём словам:")
	for _, item := range result2 {
		fmt.Printf("#%-5d %-25.25s %-55.55s...\n", item.Num, xkcdURL+"/"+strconv.Itoa(item.Num), item.Transcript)
	}
	fmt.Println("Совпадение по двум словам:")
	for _, item := range result3 {
		fmt.Printf("#%-5d %-25.25s %-55.55s...\n", item.Num, xkcdURL+"/"+strconv.Itoa(item.Num), item.Transcript)
	}
}

func alreadyInArchive(fname string) []*Comics {
	var alreadyInFile []*Comics

	rfile, err := os.OpenFile(fname, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Открыие файла на чтение: %v\n", err)
	}
	defer rfile.Close()

	if err := json.NewDecoder(rfile).Decode(&alreadyInFile); err != nil {
		log.Fatal("Ошибка декодирования:", err)
	}

	return alreadyInFile
}

// getComs - проверяет содержимое архива и при необходимости дополняет его
func getComs(num int, end int, fname string) { // добавить скачивание в несколько потоков, когда дойду до этого в учебнике =)
	alreadyInFile := alreadyInArchive(fname)

	wfile, err := os.OpenFile(fname, os.O_WRONLY, 0644) // дескриптор указывает на то что файл открыт на перезапись
	if err != nil {
		log.Fatalf("Открыие файла на чтение: %v\n", err)
	}
	defer wfile.Close()

	for num <= end {
		q := strings.Join([]string{xkcdURL, strconv.Itoa(num), "info.0.json"}, "/")
		num++
		resp, err := http.Get(q)
		if err != nil {
			fmt.Printf("Ошибка при загрузке комикса: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Статус не ОК") // комикс №404 выдаёт StatusCode 404 =)
			continue
		}

		var result Comics
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatal("Ошибка декодирования:", err)
		}

		fmt.Println(result.Num)

		alreadyInFile = append(alreadyInFile, &result)
	}

	// декодировать содержимое файла в струкруру, добавить новые данные, закодировать обратно

	marshaled, _ := json.MarshalIndent(alreadyInFile, "", " ")
	wfile.WriteString(string(marshaled) + "\n")
}

func lastIndex(fname string) int {
	file, err := os.OpenFile(fname, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []Comics
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		log.Fatalf("Ошибка маршалинга: %v\n", err)
	}

	index := result[len(result)-1].Num
	// fmt.Printf("lastIndex считает, что номер последнего комикса %d\n", index)

	return index
}

func lastOnSite() int {
	resp, err := http.Get(xkcdURL + "/" + "info.0.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result Comics
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Ошибка маршалинга: %v\n", err)
	}

	return result.Num
}
