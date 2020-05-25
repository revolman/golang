package main

import (
	"sort"
	"testing"
)

// Для проверки по каждому символу в слове
type byBytes []byte

func (b byBytes) Len() int           { return len(b) }
func (b byBytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byBytes) Less(i, j int) bool { return b[i] < b[j] }
func (b byBytes) Sort()              { sort.Sort(b) }

func TestIsPalindrome(t *testing.T) {
	// проверка числовых палиндромов
	i1 := []int{1, 2, 3, 3, 2, 1}
	if !IsPalindrome(sort.IntSlice(i1)) {
		t.Error("int: 123321 expected True")
	}

	i2 := []int{1, 2, 3, 4, 5, 6}
	if IsPalindrome(sort.IntSlice(i2)) {
		t.Error("int: 123456 expected False")
	}

	// побуквенная проверка письменных палиндромов
	s1 := "abccba"
	if !IsPalindrome(byBytes(s1)) {
		t.Errorf("byBytes: \"abccba\" expected True")
	}

	s2 := "abcdef"
	if IsPalindrome(byBytes(s2)) {
		t.Errorf("byBytes: \"abcdef\" expected False")
	}

	// проверка выражений-полиндромов
	ss1 := []string{"aaa", "bbb", "aaa"}
	if !IsPalindrome(sort.StringSlice(ss1)) {
		t.Error("StringSlice: [\"aaa\", \"bbb\", \"aaa\"] expected true")
	}

	ss2 := []string{"aaa", "bbb", "bbb"}
	if IsPalindrome(sort.StringSlice(ss2)) {
		t.Error("StringSlice: [\"aaa\", \"bbb\", \"aaa\"] expected false")
	}

}
