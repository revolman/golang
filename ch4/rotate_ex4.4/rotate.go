// Упражненеи 4.4 - версия rotate,
// которая работает в один проход
package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("rotate: %d\n", rotate(&data, 2))
}

func rotate(data *[7]int, rate int) []int {
	tmp := make([]int, rate)            // создаю массив для временного хранения значений
	copy(tmp, data[:rate])              // наполняю временный массив первыми rate-элементами data
	copy(data[:], data[rate:])          // копирую в data все элементы после rate
	copy(data[len(data)-rate:], tmp[:]) // копирую в конец data rate-элементы из временного массива
	return data[:]
}
