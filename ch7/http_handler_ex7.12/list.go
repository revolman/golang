package main

import (
	"html/template"
	"net/http"
	"os"
)

var productlist = template.Must(template.New("productList").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<title>Product List</title>
	</head>
	<body>
		<h1>Product List</h1>
		<table cellspacing="2" border="1" cellpadding="5" width="600">
			<tr>
				<th>Title</th>
				<th>Price</th>
			</tr>
			{{range $key, $value := .}}
			<tr>
				<td>{{ $key }}</td>
				<td>{{ $value }}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>
`))

// List это общая версия метода list. Выводит текущее состояние БД.
func List(w http.ResponseWriter, db database) {
	productlist.Execute(w, db)
	os.File
}
