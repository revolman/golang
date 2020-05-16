package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Параметры " + strings.Join(os.Args[0:], " "))
	fmt.Printf("Время выполнения: %f\n", time.Since(start).Seconds())
}
