// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"bytes"
	"fmt"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() (result string) {
	var sorted []int
	sorted = appendValues(sorted, t)

	b := &bytes.Buffer{}

	b.WriteByte('[')
	for i := 0; i <= len(sorted)-2; i++ {
		fmt.Fprintf(b, "%d ", sorted[i])
	}
	fmt.Fprintf(b, "%d", sorted[len(sorted)-1])
	b.WriteByte(']')

	return b.String()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add ...
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
