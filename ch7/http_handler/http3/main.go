// на случай, если нужно запутсить несколько серверов в одном приложении
// нужно создать ещё один ServeMux и выполнить ещё один вызов ListenAndServe,
// возможно параллельно.
package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float64
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	// ServeMux сопоставляет URL и обработчики, которые поделены на отдельные методы
	mux := http.NewServeMux()
	// Регистрация обработчиков
	mux.Handle("/list", http.HandlerFunc(db.list)) // db.list - значение-метод
	mux.Handle("/price", http.HandlerFunc(db.price))
	// Возможно упростить до:
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)

	// ServeMux используется в качестве основного обработчика
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "нет товара: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
