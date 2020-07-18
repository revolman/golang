// sync.Mutex

package bank3

import "sync"

var (
	mu      sync.Mutex // бинарный семафор для
	balance int        // защиты balance
)

// Deposit ...
func Deposit(amount int) {
	mu.Lock() // захват маркера
	balance += amount
	mu.Unlock() // освобождение маркера
}

// Balance ...
func Balance() int {
	mu.Lock() // захват маркера
	defer mu.Unlock()
	return balance
}
