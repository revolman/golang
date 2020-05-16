package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for index, arg := range os.Args[0:] {
		fmt.Println(index, arg)
	}
	fmt.Printf("Время выполнения: %f\n", time.Since(start).Seconds())
}
