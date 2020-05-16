package main

import (
	"encoding/json"
	"fmt"
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
	Title      string
	Transcript string
}

const xkcdURL string = "https://xkcd.com"

func main() {
	getComs()
}

func getComs() {
	var allComs []*Comics

	for num := 1; num <= 5; num++ {
		q := strings.Join([]string{xkcdURL, strconv.Itoa(num), "info.0.json"}, "/")
		resp, err := http.Get(q)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Статус не ОК")
			continue
		}

		var result Comics
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatalf("Ошибка маршалинга: %v\n", err)
		}

		allComs = append(allComs, &result)
		fmt.Println(result.Num)
	}

	// fmt.Printf("Тест: \nNum: %d\tTitle: %s\tTranscript: %s\n",
	// 	allComs[1].Num, allComs[1].Title, allComs[1].Transcript)

	file, err := os.OpenFile("comics.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// for _, item := range allComs {
	// 	fileM, _ := json.MarshalIndent(item, "", " ")
	// 	file.WriteString(string(fileM) + "\n")
	// }
	fileM, _ := json.MarshalIndent(allComs, "", " ")
	file.WriteString(string(fileM) + "\n")

}
