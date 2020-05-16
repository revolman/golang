// Server1 - минимальный "echo"-сервер.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик возвращает компонент пути из URL запроса.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
