package text_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/yakdash/pkg/modules/text"
	"github.com/stretchr/testify/assert"
)

func TestTextHasNoInit(t *testing.T) {
	m := text.New("hello world")

	cmd := m.Init()
	assert.Nil(t, cmd)
}

func TestTextDisplaysText(t *testing.T) {
	m := text.New("hello world")

	assert.Equal(t, "hello world", m.View(), "Text should be initialized to the given text")

	// Send some random message
	m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("a")})
	assert.Equal(t, "hello world", m.View(), "Text should not change on update")
}
