package panes

import tea "github.com/charmbracelet/bubbletea"

type Pane struct {
	children []Pane
	model    tea.Model
}

func NewModel(m tea.Model) Pane {
	return Pane{
		model: m,
	}
}

func (m Pane) Init() tea.Cmd {
	return nil
}

func (m Pane) Update(msg tea.Msg) (Pane, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	if m.model != nil {
		m.model, cmd = m.model.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		for i, child := range m.children {
			m.children[i], cmd = child.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Pane) View() string {
	if m.model != nil {
		return m.model.View()
	}

	return "NOT IMPLEMENTED"
}
