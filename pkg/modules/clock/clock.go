package clock

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
}

func New() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	// TODO: Ticker start
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	// TODO: Update via ticker
	return time.Now().String()
}
