package main

import "fmt"

// ByteCounter накапливает количество байтов
type ByteCounter int

// Возвращение значения в данном случае это желание соответствовать
// контракту io.Writer
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // здесь используется вывод значения C напрямую

	c = 0
	var name = "Dolly"

	// вариант получения значения через return
	i, _ := fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(i)

	c = 0
	// !!! полученчие значения напрямую после вызова метода Write через контракт с io.Writer
	fmt.Fprintf(&c, "Привет, %s", name)
	fmt.Println(c)
}
