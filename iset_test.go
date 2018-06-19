package iset

import "testing"

func TestIsEmptyWhenNothingHasBeenAdded(t *testing.T) {
	set := NewArrayIntegerSet()

	if !set.IsEmpty() {
		t.Error("Expected set to be empty but was not.")
	}
}

func TestIsEmptyWhenOneElementHasBeenAdded(t *testing.T) {
	set := NewArrayIntegerSet()

	set.Add(1)

	if set.IsEmpty() {
		t.Error("Expected set not to be empty but was.")
	}
}

func TestContains(t *testing.T) {
	set := NewArrayIntegerSet()

	set.Add(1)

	if !set.Contains(1) {
		t.Error("Expected set to contain 1 but did not.")
	}

	if set.Contains(2) {
		t.Error("Expected set not to contain 2 but did.")
	}

	set.Add(2)

	if !set.Contains(2) {
		t.Error("Expected set to contain 2 but did not.")
	}

	if !set.Contains(1) {
		t.Error("Expected set to contain 1 but did not.")
	}
}

func TestSize(t *testing.T) {
	set := NewArrayIntegerSet()

	size := set.Size()

	if size != 0 {
		t.Errorf("Expected size to be 0 but was %d", size)
	}

	set.Add(1)
	size = set.Size()

	if size != 1 {
		t.Errorf("Expected size to be 1 but was %d", size)
	}

	set.Add(2)
	size = set.Size()

	if size != 2 {
		t.Errorf("Expected size to be 2 but was %d", size)
	}

	set.Add(2)
	size = set.Size()

	if size != 2 {
		t.Errorf("Expected size to be 2 but was %d", size)
	}
}
