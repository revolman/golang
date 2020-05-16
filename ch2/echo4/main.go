// выводит аргументы командной строки
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск символа новой строки") // flag. возвращает указатель на n (*n)
var sep = flag.String("s", " ", "разделитель") // flag. возвращает указатель на sep (*sep)

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
