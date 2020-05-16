package main

import "fmt"

func main() {
	a := 5901
	b := "rsc"
	c := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbcccccccccccccccccccccccccccccc"
	fmt.Printf("#%-5d %9.9s %.55s\n", a, b, c)
}
