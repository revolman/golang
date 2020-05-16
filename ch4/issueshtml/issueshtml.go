package main

import (
	"html/template"
	"log"
	"os"

	"../github"
)

var report = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} тем</h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>Состояние</th>
		<th>Пользователь</th>
		<th>Заголовок</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
