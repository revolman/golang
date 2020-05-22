package main

import (
	"fmt"
	"sort"
)

// StringSlice тип, который будет соответствовать контракту sort.Interface
type StringSlice []string

// Len возвращает длину значения типа
func (p StringSlice) Len() int { return len(p) }

// Less обеспечивает сравнение. Если True - не менять, если False - поменять элементы
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }

// Swap представляет способ обмена
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type StringSlice1 []string

// Len возвращает длину значения типа
func (p StringSlice1) Len() int { return len(p) }

// Less обеспечивает сравнение. Если True - не менять, если False - поменять элементы
func (p StringSlice1) Less(i, j int) bool { return p[i] > p[j] }

// Swap представляет способ обмена
func (p StringSlice1) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	names := []string{"John", "Mary", "Ashley", "Rob", "Peter"}
	fmt.Println(names)
	sort.Sort(StringSlice(names))
	fmt.Println(names)
	sort.Sort(StringSlice1(names))
	fmt.Println(names)
	names = append(names, "Karl")
	fmt.Println(names)
}
