package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// обработка ключей
		// key1 := r.URL.Query().Get("cells")
		// cells, err := strconv.Atoi(key1)
		// if err != nil {
		// 	log.Printf("Ошибка cells: %v", err)
		// }
		// log.Printf("%d\n", cells)

		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
