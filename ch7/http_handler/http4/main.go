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
	// на самом деле здесь используется DefaultServeMux,
	// по этому не нужно создавать отдельеный экземпляр ServeMux.
	// В ListenAndServe нудно указвать в качестве обработчика nil
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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
