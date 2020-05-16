// Упражнение 2.5
package main

import "fmt"

// PopCount asd
func PopCount(x uint64) int {
	var result int
	for x != 0 {
		// x-1 инвертирует младшие нули, затем производится умножение.
		// Получается, что младшая 1 отбратывается. Успешная итерация записывается в результат.
		// Количество успешных итераций = количеству единичных битов.
		x &= x - 1
		result++
		// Получается, что младшая 1 отбратывается
		fmt.Printf("x = %b, result = %d\n", x, result) // Добавил для наглядности
	}
	return result
}

func main() {
	fmt.Printf("x = %b, popcount = %d\n", 33445566, PopCount(33445566))
}
