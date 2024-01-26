package panes

import (
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Pane struct {
	children []Pane
	model    tea.Model

	width  int
	height int
}

func NewModel(m tea.Model) Pane {
	return Pane{
		model: m,
	}
}

func (m Pane) WithDimensions(width, height int) Pane {
	m.width = width
	m.height = height

	return m
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

func randomColorHex() string {
	hexes := []rune("0123456789abcdef")
	color := "#"
	for i := 0; i < 6; i++ {
		color += string(hexes[rand.Intn(len(hexes))])
	}
	return color
}

func (m Pane) View() string {
	style := lipgloss.NewStyle().Width(m.width).Height(m.height).Background(lipgloss.Color(randomColorHex()))
	if m.model != nil {
		return style.Render(m.model.View())
	}

	return "NOT IMPLEMENTED"
}
