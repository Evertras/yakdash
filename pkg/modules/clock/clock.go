package clock

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	text string
}

func New() model {
	return model{
		text: curTime(),
	}
}

type tickMsg struct{}

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func curTime() string {
	return time.Now().Format("15:04:05")
}

func (m model) Init() tea.Cmd {
	return doTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		m.text = curTime()
		return m, doTick()
	}

	return m, nil
}

func (m model) View() string {
	return m.text
}
