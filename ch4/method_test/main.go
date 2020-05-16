package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const q string = "http://localhost:8000/1111"

type IssueCache struct {
	Issues         []Issue
	IssuesByNumber map[int]Issue // это отображение представляет собой кэш в виже структуры с привязкой к её номеру
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	User      *User
	Title     string
	Body      string
	State     string
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

var is1, is2 Issue
var issues []Issue

func newCache() (ic IssueCache, err error) {
	var all IssueCache
	is1.Number = 1111 // Заполняю структуры для теста >>>
	is1.Title = "TEST_TITLE"
	is2.Number = 2222
	issues = append(issues, is1)
	issues = append(issues, is2)
	all.Issues = issues          // <<< Заполнил
	mymap := make(map[int]Issue) // создаю отображение для кэша
	for _, item := range issues {
		mymap[item.Number] = item // заполняю отображение
	}
	all.IssuesByNumber = mymap // заливаю отображение в структуру
	
	ic.
	fmt.Println("is1: ", all.IssuesByNumber[1111])
	all.printer() // метод printer обрабатывает полученные данные

}

func (ic IssueCache) printer() {
	pathParts := strings.SplitN(q, "/", -1) // разбивка url на части по /
	fmt.Println(pathParts)
	numStr := pathParts[3] // элемент под этим номером значит номер issue
	fmt.Println(numStr)
	num, _ := strconv.Atoi(numStr)
	issue, _ := ic.IssuesByNumber[num] // получаем issue по взятому из url номеру
	fmt.Println(issue)                 // вывод
}
