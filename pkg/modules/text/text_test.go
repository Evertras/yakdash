package text_test

import (
	"testing"

	"github.com/evertras/yakdash/pkg/modules/text"
)

func TestTextDisplaysText(t *testing.T) {
	m := text.New("hello world")

	if m.View() != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", m.View())
	}
}
