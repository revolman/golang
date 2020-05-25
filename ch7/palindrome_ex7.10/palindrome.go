package main

import (
	"sort"
)

// IsPalindrome проверяет, является ли строка палиндромом.
// s.Less(i, j) принимает индексы элементов, которые нужно сравнить.
// Вычисляется длинна массива, заметем сравниваются первый и последний
// его элементы, затем второй и предпоследний и т.д.
// Сравнение прекращается, когда количество итераций достигает length/2.
// На каждой итерации проверяется равенство первого и последнего элемента функцией:
// !s.Less(i, j) && !s.Less(j, i)
// т.к. s.Less возвращает true только когда i < j равенством элементов считается
// !(i<j) && !(j<i)
// Пример: s := "A B C C B A"; len(s) = 6;
// i=0: !s.Less(s[0], s[5]) && !s.Less(s[5], s[0]) --> true
func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i := 0; i < length/2; i++ {
		if !s.Less(i, length-i-1) && !s.Less(length-i-1, i) {
			continue
		} else {
			return false
		}
	}
	return true
}
