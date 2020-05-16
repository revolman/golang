package main

import "fmt"

// Path ...
type Path []string

func main() {
	perim := Path{"one", "two", "three", "one"}
	for i := range perim {
		if i > 0 {
			fmt.Printf("i:%d\tPath: %s\n", i, perim[i-1])
		}
		fmt.Println(i)
	}
}
