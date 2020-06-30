package main

import (
	"fmt"
)

var deposits = make(chan int)
var balances = make(chan int)

// Deposit - пополнение счёта
func Deposit(amount int) { deposits <- amount }

// Balance - вывод баланса
func Balance() int { return <-balances }

func teller() {
	var balance int // ограничен теллером
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func main() {
	go teller() // запуск управляющей goroutine
	Deposit(50)
	Deposit(100)
	fmt.Println(Balance())
	Deposit(200)
	fmt.Println(Balance())
}
