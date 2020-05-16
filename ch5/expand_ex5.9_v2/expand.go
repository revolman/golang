package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "Это тест $foo $foo $foo \n"

	f := func(s string) string {
		subString := strings.Join(os.Args[1:], " ")
		result := strings.Replace(s, "$foo", subString, -1)
		return result
	}
	result := expand(s, f)
	fmt.Println(result)
}

// что-то пошло не так =)
func expand(s string, f func(string) string) string {
	return f(s)
}
