package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// ReposAPI - ссылка на API
const ReposAPI = "https://api.github.com/repos/"

// IssuesURL - адрес API для работы с поиском issues
const IssuesURL = "https://api.github.com/search/issues"

var personalAccessToken = "ac44681d20e5a97ce7c4fee42719e4d35d374f67"

// IssuesSearchResult - структура для хранения результата выполненного поиска
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue - структура каждой темы
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	User      *User
	Title     string
	Body      string
	State     string
	CreatedAt time.Time `json:"created_at"`
}

// CacheURL - метод типа Issue, который подготавливает url для каждой темы.
// Из этого url будет выбираться номер темы.
// По номеру будет выводиться шаблон темы, ранее записанный в отображение
func (i Issue) CacheURL() string {
	return fmt.Sprintf("/issue/%d", i.Number)
}

// User - структура хранения инфы о пользователе
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// Comment - структура для хранения комментариев к теме
type Comment struct {
	ID        int
	Body      string
	User      *User
	CreatedAt time.Time `json:"created_at"`
}

// TokenSource - храненеи токена аутентификации
type TokenSource struct {
	AccessToken string
}

// GetComments получает список комментариев к теме
func GetComments(owner string, repo string, number string) ([]Comment, error) {
	q := strings.Join([]string{owner, repo, "issues", number, "comments"}, "/")
	resp, err := http.Get(ReposAPI + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result []Comment
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetIssues делает запрос к API и возвращает срез всех Issues
func GetIssues(owner string, repo string) ([]Issue, error) {
	q := strings.Join([]string{owner, repo}, "/")
	resp, err := http.Get(ReposAPI + q + "/issues")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// Token auth
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
