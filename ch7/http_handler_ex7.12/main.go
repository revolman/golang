package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mux sync.Mutex

type dollars float64
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {

	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/remove", db.remove)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	List(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		List(w, db)
		fmt.Fprintf(w, "нет товара: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	sprice := req.URL.Query().Get("price")

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		List(w, db)
		fmt.Fprintf(w, "нет товара: %q\n", item)
		return
	}

	iprice, err := strconv.ParseFloat(sprice, 8)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "некорректный ввод: %v\n", err)
		fmt.Fprint(w, "принимается только положительные значения с плавающей точкой, пример: 3.14")
		return
	}

	if iprice < 0 {
		List(w, db)
		fmt.Fprintf(w, "цена не может быть отрицательной: %.2f\n", iprice)
		return
	}

	mux.Lock()
	db[item] = dollars(iprice)
	mux.Unlock()
	List(w, db)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	sprice := req.URL.Query().Get("price")

	// iprice, err := strconv.Atoi(sprice)
	iprice, err := strconv.ParseFloat(sprice, 8)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "некорректный ввод: %v\n", err)
		return
	}

	if iprice < 0 {
		List(w, db)
		fmt.Fprintf(w, "цена не может быть отрицательной: %.2f\n", iprice)
		return
	}

	_, ok := db[item]
	if !ok {
		mux.Lock()
		db[item] = dollars(iprice)
		mux.Unlock()
		w.WriteHeader(http.StatusOK)
		List(w, db)
		return
	}
	List(w, db)
	fmt.Fprintf(w, "товар %q уже добавлен в базу, для изменения используйте /update\n", item)
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		List(w, db)
		fmt.Fprintf(w, "товара нет в базе: %q", item)
		return
	}
	mux.Lock()
	delete(db, item)
	mux.Unlock()
	List(w, db)
}
