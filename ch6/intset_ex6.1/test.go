package intset

import "testing"

// TestAddHas ...
func TestAddHas(t *testing.T) {
	set := new(IntSet)

	set.Add(1)
	if !set.Has(1) {
		t.Errorf("Set %s should have 1", set)
	}

	set.Add(1024)
	if !set.Has(1) && !set.Has(1024) {
		t.Errorf("Set %s should have 1 and 1024", set)
	}

	set.Add(0)
	set.Add(1024)
	if !set.Has(0) && !set.Has(1) && !set.Has(1024) {
		t.Errorf("Set %s should have 0, 1 and 1024", set)
	}

	set.Add(-1)
	if set.Has(-1) || set.Len() != 3 {
		t.Errorf("Set cannot add negative numbers. %s should have 0, 1, 1024", set)
	}
}

// TestRemove ...
func TestRemove(t *testing.T) {
	set := new(IntSet)

	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(0)
	set.Add(1024)

	set.Remove(2)
	if set.Has(2) || set.Len() != 4 {
		t.Errorf("Set %s should have 0, 1, 3 and 1024", set)
	}

	set.Remove(-1)
	if set.Len() != 4 {
		t.Errorf("Set %s should have 0, 1, 3 and 1024", set)
	}

	set.Remove(1024)
	if set.Has(1024) || set.Len() != 3 {
		t.Errorf("Set %s should have 0, 1 and 3", set)
	}

	set.Remove(0)
	if set.Has(0) || set.Len() != 2 {
		t.Errorf("Set %s should have 1 and 3", set)
	}

	set.Remove(0)
	if set.Has(0) || set.Len() != 2 {
		t.Errorf("Set %s should have 1 and 3", set)
	}

	set.Remove(1)
	set.Remove(3)
	if set.Has(1) || set.Has(3) || set.Len() != 0 {
		t.Errorf("Set %s should be empty", set)
	}

}

// TestLen ...
func TestLen(t *testing.T) {
	set := new(IntSet)
	if l := set.Len(); l != 0 {
		t.Errorf("Empty set should have o lenght. Returned: %d", l)
	}

	set.Add(1)
	set.Add(2)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}

	set.Add(2)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}

	set.Add(129)
	if l := set.Len(); l != 3 {
		t.Errorf("Set %s should have o Len() == 3. Returned: %d", set, l)
	}

	set.Remove(2)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}

	set.Remove(3)
	if l := set.Len(); l != 2 {
		t.Errorf("Set %s should have o Len() == 2. Returned: %d", set, l)
	}
}

// TestClear ...
func TestClear(t *testing.T) {
	set := new(IntSet)
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Clear()
	if set.Len() != 0 {
		t.Errorf("Set len must be 0")
	}
}
