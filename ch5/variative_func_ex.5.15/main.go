package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers []int
	args := os.Args[1:]
	for _, number := range args {
		conv, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, conv)
	}

	fmt.Println(max(numbers...))
	fmt.Println(min(numbers...))

}

func max(numbers ...int) (max int) {
	if len(numbers) == 0 {
		log.Fatal("не указаны аргументы")
	}

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return
}

// 1 3 8 4 5
func min(numbers ...int) (min int) {
	if len(numbers) == 0 {
		log.Fatal("не указаны аргументы")
	}

	min = numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return
}
