package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
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
	Items      []*Issues
}

// Issues - структура каждой темы
type Issues struct {
	Number    int
	User      *User
	Title     string
	Body      string
	State     string
	CreatedAt time.Time `json:"created_at"`
}

// User - структура хранения инфы о пользователе
type User struct {
	Login string
}

// TokenSource - храненеи токена аутентификации
type TokenSource struct {
	AccessToken string
}

// SearchAnIssues - делает http.Get запрос к API issues Github,
// записывает результат в структуру IssuesSearchResult
// и возвращает *IssuesSearchResult или error
func SearchAnIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnIssues делает запрос к API и возвращает срез всех Issues
func GetAnIssues(owner string, repo string) ([]Issues, error) {
	q := strings.Join([]string{owner, repo}, "/")
	resp, err := http.Get(ReposAPI + q + "/issues")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result []Issues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateAnIssue универсальная функция обновления issue
func UpdateAnIssue(method string, query string, data map[string]string) (*Issues, error) {
	// преобразование данных в формат JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(jsonData)

	// В комментарии указан Depricated способо аутентификации
	// old auth >
	// client := &http.Client{} // подготовка клиента
	// username, password := credentials()
	// req, err := http.NewRequest(method, query, buf)
	// req.SetBasicAuth(username, password)
	// if err != nil {
	// 	return nil, err
	// }
	// resp, err := client.Do(req) // клиент делает запрос, получает ответ или ошибку
	// < old auth

	req, err := http.NewRequest(method, query, buf)
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	resp, err := oauthClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Ошибка: %v", err)
	}
	if method == "POST" && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Статус: %s", resp.Status)
	}
	if method == "PATCH" && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Статус: %s", resp.Status)
	}
	defer resp.Body.Close()

	var result Issues

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAnIssue возвращает тему, которую можно использовать, к примеру для редактирования
func GetAnIssue(owner string, repo string, number string) (*Issues, error) {
	q := strings.Join([]string{owner, repo, "issues", number}, "/")
	resp, err := http.Get(ReposAPI + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result Issues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Получаю username и password
// Для скрытого ввода пароля используется библиотека golang.org/x/crypto/ssh/terminal.
// Аутентификация по логину и паролю признана устаревшей, данная функция не используется.
// Заменил на oauth2, для аутентификации через токен.
// Оставлено в коде как памятка.
func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите логин: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nОшибка: %v\n", err)
	}
	fmt.Print("Введите пароль: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nОшибка: %v\n", err)
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

// Token auth
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
