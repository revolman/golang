// Package convertator. Упражнение 2.2 - собирает в себе пакеты,
// конвертирующие значения различных единиц измерения.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../lengthconv"
	"../tempconv"
	"../weightconv"
)

var ctof = flag.Float64("cf", 0, "перевод цельсия в фаренгейт")

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Что будем конвертировать?\nВарианты: 1 Градусы, 2 Вес, 3 Длина.\n")
		choise := choiser(os.Stdin)

		// Конвертация шкал градусов
		if strings.TrimRight(choise, "\n") == "1" || strings.TrimRight(choise, "\n") == "Градусы" {
			fmt.Printf("Варинты: CF, CK, FC, FK, KC, KF\nВыбор: ")
			choise := choiser(os.Stdin)

			if strings.TrimRight(choise, "\n") == "CF" {
				fmt.Println("Цельсий в Фаренгейт: ")
				choise := choiser(os.Stdin)
				degree := tempconv.Celsius(converter(choise))
				fmt.Printf("%s = %s\n", degree, tempconv.CToF(degree))
				os.Exit(0)
			}
			if strings.TrimRight(choise, "\n") == "CK" {
				fmt.Println("Цельсий в Кельвин: ")
				choise := choiser(os.Stdin)
				degree := tempconv.Celsius(converter(choise))
				fmt.Printf("%s = %s\n", degree, tempconv.CToK(degree))
				os.Exit(0)
			}
			if strings.TrimRight(choise, "\n") == "FC" {
				fmt.Println("Фаренгейт в Цельсий: ")
				choise := choiser(os.Stdin)
				degree := tempconv.Fahrenheit(converter(choise))
				fmt.Printf("%s = %s\n", degree, tempconv.FToC(degree))
				os.Exit(0)
			}
			if strings.TrimRight(choise, "\n") == "FK" {
				fmt.Println("Фаренгейт в Цельсий: ")
				choise := choiser(os.Stdin)
				degree := tempconv.Fahrenheit(converter(choise))
				fmt.Printf("%s = %s\n", degree, tempconv.FToK(degree))
				os.Exit(0)
			}
			if strings.TrimRight(choise, "\n") == "KC" {
				fmt.Println("Кельвин в Цельсий: ")
				choise := choiser(os.Stdin)
				degree := tempconv.Kelvin(converter(choise))
				fmt.Printf("%s = %s\n", degree, tempconv.KToC(degree))
				os.Exit(0)
			}
			if strings.TrimRight(choise, "\n") == "KF" {
				fmt.Println("Кельвин в Фаренгейт: ")
				choise := choiser(os.Stdin)
				degree := tempconv.Kelvin(converter(choise))
				fmt.Printf("%s = %s\n", degree, tempconv.KToF(degree))
				os.Exit(0)
			}
		}

		// Конвертация единиц веса
		if strings.TrimRight(choise, "\n") == "2" || strings.TrimRight(choise, "\n") == "Вес" {
			fmt.Printf("Варинты: 1 фунт->килограмм, 2 килограмм->фунт\nВыбор: ")
			choise := choiser(os.Stdin)

			if strings.TrimRight(choise, "\n") == "1" {
				fmt.Println("Введите вес в фунтах: ")
				choise := choiser(os.Stdin)
				weight := weightconv.Pound(converter(choise))
				fmt.Printf("%s = %s\n", weight, weightconv.PToK(weight))
				os.Exit(0)
			}
			if strings.TrimRight(choise, "\n") == "2" {
				fmt.Println("Введите вес в килограммах: ")
				choise := choiser(os.Stdin)
				weight := weightconv.Kilogram(converter(choise))
				fmt.Printf("%s = %s\n", weight, weightconv.KToP(weight))
				os.Exit(0)
			}
		}

		// Конвертация единиц длины
		if strings.TrimRight(choise, "\n") == "3" || strings.TrimRight(choise, "\n") == "Длина" {
			fmt.Printf("Варинты: 1 фут->метр, 2 метр->фут\nВыбор: ")
			choise := choiser(os.Stdin)

			if strings.TrimRight(choise, "\n") == "1" {
				choise := choiser(os.Stdin)
				length := lengthconv.Foot(converter(choise))
				fmt.Printf("%s = %s\n", length, lengthconv.FToM(length))
				os.Exit(0)
			}
			if strings.TrimRight(choise, "\n") == "2" {
				choise := choiser(os.Stdin)
				length := lengthconv.Meter(converter(choise))
				fmt.Printf("%s = %s\n", length, lengthconv.MToF(length))
				os.Exit(0)
			}
		}
	} else {
		// Пример обработки флагов. Только cf, что бы не раздувать программу.
		// Впрочем её нверняка можно оптимизировать
		flag.Parse()
		// Костыль, который притворяется защитой от дурака.
		// Нужно разобраться, как проверять установлен ли флаг.
		if *ctof > -273.16 {
			degree := tempconv.Celsius(*ctof)
			fmt.Printf("%s = %s\n", degree, tempconv.CToF(degree))
			os.Exit(0)
		} else {
			fmt.Println("Температура ниже абсолютного нуля не возможна.")
			os.Exit(0)
		}
	}
}

func choiser(input *os.File) string {
	reader := bufio.NewReader(input)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Возникла ошибка: %v", err)
	}
	return text
}

func converter(s string) float64 {
	s = strings.TrimRight(s, "\n")
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Ошибка, ожидалось введение числа: %v", err)
	}
	return i
}
