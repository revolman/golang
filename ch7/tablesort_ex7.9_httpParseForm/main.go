// Разобрал реализацию renatofq. Интересный алгоритм.
// Но судя по всему можно было бы сделать тупее и проще.
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

var tracklist = template.Must(template.New("tracklist").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<title>Track List</title>
	</head>
	<body>
		<h1>Track List</h1>
		<table>
			<tr>
				<th><a href='/?sortBy=Title'>Title</a></th>
				<th><a href='/?sortBy=Artist'>Album</a></th>
				<th><a href='/?sortBy=Album'>Album</a></th>
				<th><a href='/?sortBy=Year'>Album</a></th>
				<th><a href='/?sortBy=Length'>Album</a></th>
			</tr>
			{{range .}}
			<tr>
				<td>{{.Title}}</td>
				<td>{{.Artist}}</td>
				<td>{{.Album}}</td>
				<td>{{.Year}}</td>
				<td>{{.Length}}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>
`))

type App struct {
	ts     *trackSort
	tracks []*Track
}

func main() {
	app := &App{ts: defLessFuncs(tracks), tracks: tracks}
	http.HandleFunc("/", app.tracklistHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (app *App) tracklistHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(400), 400)
	}

	if sortBy := r.Form.Get("sortBy"); sortBy != "" {
		app.ts.SortBy(sortBy)
		sort.Sort(app.ts)
	}
	tracklist.Execute(w, app.tracks) // передать в шаблон структуру
}
