package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Параметры " + strings.Join(os.Args[0:], " "))
}
