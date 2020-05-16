package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var issuesTemplate = template.Must(template.New("issuelist").Parse(`
	<h1>{{.Issues | len}} тем</h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>Состояние</th>
		<th>Пользователь</th>
		<th>Заголовок</th>
	</tr>
	{{range .Issues}}
	<tr>
		<td><a href='{{.CacheURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.CacheURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
`))

var issueTemplate = template.Must(template.New("issueinfo").Parse(`
	<h1>{{.Title}}</h1>
	<samll>by <a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></small>
	<p>{{.Body}}</p>
	<br>
	<h3>Коментарии:</h3>
`))

var commentsTemplate = template.Must(template.New("comments").Parse(`
	<p>#{{.ID}}, {{.User.Login}}</p>
	{{.Body}}
	<br>
`))

// IssuesCache - структура, в которую сохраняются все полученные темы
// позволяет вытягивать темы по номеру, да создания кэша тем
type IssuesCache struct {
	Issues         []Issue
	IssuesByNumber map[int]Issue
	Comments       map[int][]Comment
}

// type IssuesCache struct {
// 	Issues         []Issue
// 	IssuesByNumber map[int]Issue
// 	Comments	[]Comment
// }

// Метод ic типа IssueCache позволяет при любом вызове объекта типа IssueCache запускать функцию ServeHTTP,
// которая является обработчиком запросов.
func (ic IssuesCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := strings.SplitN(r.URL.Path, "/", -1)
	if len(query) < 3 || query[2] == "" {
		// если меньше 3х элементов или номер issue пустой
		// Тогда просто вывести весь список полученных тем по шаблону issueListTemplate
		if err := issuesTemplate.Execute(w, ic); err != nil {
			log.Print(err)
		}
		return
	}
	// Во всех остальных случаях:
	numStr := query[2]               // берётся номер темы из url
	num, err := strconv.Atoi(numStr) // конвертируется в строку
	if err != nil {                  // обработка неверных номеров
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Номер задан не числом, вот этой фигнёй: '%s'", numStr)))
		if err != nil {
			log.Printf("Не удалось создать запрос для %v: %v", r, err)
		}
		return
	}
	issue, ok := ic.IssuesByNumber[num]
	// из полученного отображения по номеру выбирается тема
	// Если тема не существует тогда обработать ошибку
	if !ok {
		_, err := w.Write([]byte(fmt.Sprintf("Не существует темы с номером %d", num)))
		if err != nil {
			log.Printf("Не удалось создать запрос для %v: %v", r, err)
		}
		return
	}
	// записать выбранную тему в ResponseWriter в виде шаблона issueTemplate
	if err := issueTemplate.Execute(w, issue); err != nil {
		log.Print(err)
	}

	comments, ok := ic.Comments[num]
	if !ok {
		_, err := w.Write([]byte(fmt.Sprintf("Не существует темы с номером %d", num)))
		if err != nil {
			log.Printf("Не удалось создать запрос для %v: %v", r, err)
		}
		return
	}

	// для каждого шаблона в списке шаблонов делать запись во врайтер
	for _, comment := range comments {
		if err := commentsTemplate.Execute(w, comment); err != nil {
			log.Printf("Преобразование шаблона commentsTemplate: %v", err)
		}
	}

}

// getNewCache Получает все Issues в указанном репозитории
// Результат: структура IssueCache.
// В результат записывает полученные данные. Использует метод ic.
// ic IssueCache равно var ic IssueCache, только в виде именованного параметра функции
func getNewCache(owner string, repo string) (ic IssuesCache, err error) {
	issues, err := GetIssues(owner, repo)
	if err != nil {
		log.Fatalf("Ошибка при получении Issues: %v", err)
	}

	ic.Issues = issues
	// Создаю структуры по поличеству issues
	ic.IssuesByNumber = make(map[int]Issue, len(issues))
	ic.Comments = make(map[int][]Comment, len(issues))
	for _, issue := range issues {
		ic.IssuesByNumber[issue.Number] = issue
		// Добавлять список комментов для каждой issue в отдельное поле IC
		numStr := strconv.Itoa(issue.Number)
		ic.Comments[issue.Number], err = GetComments(owner, repo, numStr)
		if err != nil {
			log.Fatalf("Ошибка при получении комментариев: %v", err)
		}
	}

	return
}

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Обязательно нужно указать аргументы: owner repo")
		os.Exit(0)
	}

	owner := os.Args[1]
	repo := os.Args[2]

	issuesCache, err := getNewCache(owner, repo)
	if err != nil {
		log.Print(err)
	}

	http.Handle("/", issuesCache)                         // вызов типа IssueCache запускает метод ServeHTTP
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // сервер
}
