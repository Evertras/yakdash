package id

import "testing"

func TestNew(t *testing.T) {
	saw := map[ID]bool{}

	for i := 0; i < 1000; i++ {
		id := New()

		if saw[id] {
			t.Errorf("Saw duplicate ID %d", id)
		}

		saw[id] = true
	}
}
