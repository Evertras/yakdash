package id_test

import (
	"testing"

	"github.com/evertras/yakdash/pkg/id"
)

func TestNew(t *testing.T) {
	saw := map[id.ID]bool{}

	for i := 0; i < 1000; i++ {
		id := id.New()

		if saw[id] {
			t.Errorf("Saw duplicate ID %d", id)
		}

		saw[id] = true
	}
}
