package main

import (
	"fmt"
)

// Wthd ...
type Wthd struct {
	amount int         // сумма текущего снятия
	status chan<- bool // передаю канал в соответствии с условиями задачи
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan Wthd)

// Deposit - пополнение счёта
func Deposit(amount int) { deposits <- amount }

// Withdraw - снятие со счёта
func Withdraw(amount int) bool {
	var check = make(chan bool)      // создаю канал статусов для проверки хватает ли денег
	withdraws <- Wthd{amount, check} // передаю сумму, канал передаётся для проверки хватает ли денег
	return <-check                   // возвращаю значение проверки, можно использовать его для вывода
	// предупреждения о нехватке средств
}

// Balance - вывод баланса
func Balance() int { return <-balances }

func teller() {
	var balance int // ограничен теллером
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case wthMsg := <-withdraws:
			ok := balance >= wthMsg.amount // нельзя снять больше чем есть
			if ok {                        // если на счёте хватает денег, то можно их снять с баланса
				balance -= wthMsg.amount
			}
			wthMsg.status <- ok // передаю результат проверки в check
		case balances <- balance: // обновляю баланс
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
	if Withdraw(33) {
		fmt.Println("OK!")
	} else {
		fmt.Println("Not enough money!")
	}
	fmt.Println(Balance())
	if Withdraw(500) {
		fmt.Println("OK!")
	} else {
		fmt.Println("Not enough money!")
	}
	fmt.Println(Balance())
}
