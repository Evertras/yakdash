package text

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	text string
}

func New(text string) model {
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
