// Echo1 выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ", "
	}
	fmt.Println("Аргументы: " + s)
}
