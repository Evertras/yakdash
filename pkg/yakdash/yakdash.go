package yakdash

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/evertras/yakdash/pkg/layout"
	"github.com/evertras/yakdash/pkg/panes"
)

type model struct {
	layout layout.Root

	rootPane panes.Pane
}

func (m model) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)

	cmds = append(cmds, tea.EnterAltScreen)
	cmds = append(cmds, m.rootPane.Init())

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.rootPane, cmd = m.rootPane.WithDimensions(msg.Width, msg.Height)
		cmds = append(cmds, cmd)
	}

	m.rootPane, cmd = m.rootPane.Update(msg)

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.rootPane.View()
}
