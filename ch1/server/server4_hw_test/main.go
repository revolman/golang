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
	keys, ok := r.URL.Query()["cycles"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
		return
	}
	cycles := keys[0]
	fmt.Fprintf(w, "cycles = %v\n", cycles)
}
