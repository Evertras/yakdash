package text_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/yakdash/pkg/modules/text"
	"github.com/stretchr/testify/assert"
)

func TestTextHasNoInit(t *testing.T) {
	m := text.NewPlainText("hello world")

	cmd := m.Init()
	assert.Nil(t, cmd)
}

func TestTextDisplaysText(t *testing.T) {
	m := text.NewPlainText("hello world")

	assert.Equal(t, "hello world", m.View(), "Text should be initialized to the given text")

	// Send some random message
	m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("a")})
	assert.Equal(t, "hello world", m.View(), "Text should not change on update")
}

func TestTextParsesConfig(t *testing.T) {
	cfg := map[string]interface{}{
		"text": "hello",
	}

	m, err := text.New(cfg)

	assert.NoError(t, err)

	assert.Equal(t, "hello", m.View())
}
