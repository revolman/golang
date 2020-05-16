package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // Не забывайте о скобках!
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Вход в %s", msg)
	return func() { log.Printf("Выход из %s (%s)", msg, time.Since(start)) }
}
