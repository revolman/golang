package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
)

// APIKey ...
const APIKey = "b414d24b"

// APIURL ...
const APIURL = "http://www.omdbapi.com/"

// тест шаблонов >>>>
const searchTempl = `{{.TotalResults}} результатов:
{{range .Search}}---------------------------------
Постер: {{.Poster}}
Название: {{.Title}}
Год выпуска: {{.Year}}
Тип: {{.Type}}
{{end}}`

// Movie ...
type Movie struct {
	Title      string
	Year       string
	Released   string
	Director   string
	Country    string
	Poster     string
	Ratings    []*Ratings
	IMDBRating string
	Type       string
	Actors     string
	Plot       string
	ImdbID     string
	Response   string
}

// Ratings ...
type Ratings struct {
	Source string
	Value  string
}

// SearchingResult хранит все результатs поиска, но нас сейчас интересует только первый фильм
type SearchingResult struct {
	Search       []*Search
	TotalResults string
	Response     string
}

// Search ...
type Search struct {
	Title  string
	Year   string
	imdbID string
	Type   string
	Poster string
}

// тест шаблонов >>>>
var searchReport = template.Must(template.New("searchReport").
	Parse(searchTempl))

func main() {
	if len(os.Args[1:]) < 1 {
		Usage()
	}

	cmd := os.Args[1]
	query := strings.Join(os.Args[2:], " ")

	switch cmd {
	case "search":
		searchResult, err := Searching(query)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		// тест шаблонов >>>>
		if err := searchReport.Execute(os.Stdout, searchResult); err != nil {
			log.Fatal(err)
		}

		// fmt.Println("Всего найдено: ", searchResult.TotalResults)
		// for _, item := range searchResult.Search {
		// 	fmt.Printf("%s - %s, %s\n", item.Title, item.Year, item.Type)
		// }
	case "poster":
		// вывод афиши >>>>
		poster, err := GetPoster(query)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		if poster.Response == "True" {
			fmt.Println("Информация о фильме ", poster.Title)
			fmt.Printf("\nПостер: %s\n\n"+
				"Дата выпуска: %s\n"+
				"Режиссёр: %s\n"+
				"В ролях: %s\n"+
				"Рейтинг IMDB: %s\n"+
				"\nСюжет: %s\n\n", poster.Poster, poster.Released, poster.Director, poster.Actors, poster.IMDBRating, poster.Plot)
		} else {
			fmt.Println("К сожалению, фильма с таким названием не найдено.")
		}
	default:
		Usage()
	}
}

// Searching - получает список фильмов в соответствии с запросом
// поиск по заданию не нужен.
func Searching(query string) (*SearchingResult, error) {
	q := url.QueryEscape(query)
	resp, err := http.Get(APIURL + "?apikey=" + APIKey + "&s=" + q)
	if err != nil {
		return nil, fmt.Errorf("searching: Не удаётся установить соединение")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result SearchingResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Ошибка маршалинга: %v", err)
	}

	return &result, nil
}

// GetPoster получает информацию о конкретном фильме
func GetPoster(query string) (*Movie, error) {
	q := url.QueryEscape(query)
	resp, err := http.Get(APIURL + "?apikey=" + APIKey + "&t=" + q)
	if err != nil {
		return nil, fmt.Errorf("getPoster: Не удалось установить соединение")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Ошибка маршалинга %v", err)
	}

	return &result, nil
}

// Usage завершает программу и выводит подсказку по использованию
func Usage() {
	fmt.Printf("Использование:\nsearch [query] - поиск фильмов по названию\n" +
		"poster [film name] - вывод афиши фильма\n")
	os.Exit(1)
}
