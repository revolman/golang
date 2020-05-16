// Используя panic и recover написать функцию, которая не содержит инструкции return,
// но возвращает не нулевое значение.
// В моём случае программа паникует и возвращает ошибку в случае получения любых мастей карт,
// которые отличаются от указанных в кейсах.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	suit := strings.Title(os.Args[1])
	err := cards(suit)
	fmt.Println(err)
}

func cards(suit string) (err error) {
	type bailout struct{}

	fmt.Printf("Ввод: %s\n", suit)
	defer func() {
		switch p := recover(); p {
		case nil:
			// не паникую
		case bailout{}:
			// ожидаю панику, по этому готовлю восстановление
			err = fmt.Errorf("English motherfucker, do you speak it?")
			// panic(fmt.Sprintf("English motherfucker, do you speak it?\n", suit))
		default:
			// не ожидаю - паникую
			panic(fmt.Sprintf("неверная карта %q", suit))
		}
	}()

	switch suit {
	case "Spades":
		fmt.Println("OK")
		os.Exit(0)
	case "Hearts":
		fmt.Println("OK")
		os.Exit(0)
	case "Diamonds":
		fmt.Println("OK")
		os.Exit(0)
	case "Clubs":
		fmt.Println("OK")
		os.Exit(0)
	}
	panic(bailout{})
}
