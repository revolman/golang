package treesort

import "testing"

func TestString(t *testing.T) {
	root := &tree{value: 7}
	root = add(root, 3)
	root = add(root, 5)
	if root.String() != "[3 5 7]" {
		t.Errorf("Wrong string in result: %s", root.String())
	}
}
