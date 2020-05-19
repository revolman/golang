package bytecounter

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var c WordCounter
	var name = "Dolly"

	fmt.Fprintf(&c, "Hello, %s! How are you?", name)
	if c != 5 {
		t.Errorf("c must be 5, not %d", c)
	}
}

func TestLineCounter(t *testing.T) {
	var l LineCounter

	fmt.Fprintf(&l, "Line1\nLine2\nLine3\n")

	if l != 3 {
		t.Errorf("l must be 3, not %d", l)
	}
}
