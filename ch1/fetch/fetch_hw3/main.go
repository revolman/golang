// Упражнение 1.9.
// Вывод кода состояния HTTP.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http") != true { // Проверка префикса http*
			url = "https://" + url // Добаление префикса https://
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", resp.Status)
	}
}
