// sync.Mutex

package bank2

var (
	sema    = make(chan struct{}, 1) // бинарный семафор для
	balance int                      // защиты balance
)

// Deposit ...
func Deposit(amount int) {
	sema <- struct{}{} // захват маркера
	balance += amount
	<-sema // освобождение маркера
}

// Balance ...
func Balance() int {
	sema <- struct{}{} // захват маркера
	b := balance       // передача значения переменной
	<-sema             // освобождение маркера
	return b
}
