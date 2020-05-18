package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

const wordSize = 32 << (^uint(0) >> 63)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Len ...
func (s *IntSet) Len() int {
	var count uint64
	for _, word := range s.words {
		for word != 0 {
			count += word & 1
			word = word >> 1
		}
	}
	return int(count)
}

// Clear ...
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy ...
func (s *IntSet) Copy() *IntSet {
	c := new(IntSet)
	// for _, word := range s.words {
	// 	c.words = append(c.words, word)
	// }
	c.words = make([]uint64, len(s.words))
	copy(c.words, s.words)
	return c
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll Add-like variative method
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Remove ...
func (s *IntSet) Remove(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith Поиск пересечения в двух срезах слов
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range t.words {
		if i < len(s.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] &= 0 // на случай, если в одном срезе больше слов, чем в другом
		}
	}
}

// SymmetricDifferenceWith ...
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword // XOR
		} else {
			// т.к. это симметричная разница, то все элементы у которых нет пары для операции нужно просто добать в результат
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith ...
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range t.words {
		if i < len(s.words) {
			s.words[i] &^= t.words[i] // штрих шейфера (и-не). вопрос к нулям разве что...
		} else {
			break // здесь не нужно обрабатывать этементы без пары
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Elems возвращет срез элементов множества
func (s *IntSet) Elems() []int {
	var e []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint64(j)) != 0 {
				e = append(e, wordSize*i+j)
			}
		}
	}
	return e
}

//!-string
