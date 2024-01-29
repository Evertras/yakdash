package text

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mitchellh/mapstructure"
)

type model struct {
	text string
}

func New(cfg map[string]interface{}) (model, error) {
	var config config
	err := mapstructure.Decode(cfg, &config)

	if err != nil {
		return model{}, fmt.Errorf("failed to parse config: %w", err)
	}

	return model{
		text: config.Text,
	}, nil
}

func NewPlainText(text string) model {
	return model{
		text: text,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return m.text
}
