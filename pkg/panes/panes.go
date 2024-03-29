package panes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Direction int

const (
	DirectionVertical Direction = iota
	DirectionHorizontal
)

type Pane struct {
	children  []Pane
	model     tea.Model
	direction Direction

	name string

	style lipgloss.Style

	width  int
	height int
}

// NewModel creates a new pane containing the given model.
func NewLeaf(m tea.Model) Pane {
	borderStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

	return Pane{
		model: m,
		style: borderStyle.Align(lipgloss.Top, lipgloss.Left),
	}
}

// NewNode creates a new pane containing the given children.
func NewNode(direction Direction, children ...Pane) Pane {
	return Pane{
		direction: direction,
		children:  children,
	}
}

func (p Pane) WithName(name string) Pane {
	p.name = name
	p.style = p.style.BorderTop(p.name == "")

	return p
}

func (m Pane) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)

	if m.model != nil {
		cmds = append(cmds, m.model.Init())
	} else {
		for _, child := range m.children {
			cmds = append(cmds, child.Init())
		}
	}

	return tea.Batch(cmds...)
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
