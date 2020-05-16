package main

import (
	"fmt"
	"os"
	"strconv"

	"../tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка cf: %v", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		fmt.Printf("C -> F: %s = %s\nC -> K: %s = %s\nF -> C: %s = %s\nF -> K: %s = %s\nK -> C: %s = %s\nK -> F: %s = %s\n",
			c, tempconv.CToF(c), c, tempconv.CToK(c),
			f, tempconv.FToC(f), f, tempconv.FToK(f),
			k, tempconv.KToC(k), k, tempconv.KToF(k))
	}
}
