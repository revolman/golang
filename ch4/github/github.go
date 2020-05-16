package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL - адрес API для работы с issues
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult - структура для хранения результата выполненного поиска
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issues
}

// Issues - структура для хранения полученных issues
type Issues struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Title     string
}

// User - структура для хранения данных о создателях issues
type User struct {
	HTMLURL string `json:"html_url"`
	Login   string
}

// SearchIssues - делает http.Get запрос к API issues Github,
// записывает результат в структуру IssuesSearchResult
// и возвращает *IssuesSearchResult или error
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}
