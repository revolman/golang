// обработчик HTTP-запросов
// обработка ключей, полученных в url. Пример /?cycles=20.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["cycles"] // извлечение из url значение ключа cycles
		if !ok || len(keys[0]) < 1 {        // если ключ не получен, то вывод предупреждения
			fmt.Fprintf(w, "Параметр cycles не указан.")
			return
		}
		cycles, err := strconv.Atoi(keys[0]) // конвертирование значения ключа в int
		if err != nil {                      // защита от дурака
			fmt.Fprintf(w, "Параметр cycles должен быть целым числом: %s", err)
			return
		}
		lissajous(w, cycles) // передача значения в ответ
	}
	http.HandleFunc("/", handler)                         // обработчик
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // сервер
}
