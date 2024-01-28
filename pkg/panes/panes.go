package panes

import (
	"github.com/charmbracelet/bubbles/viewport"
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

	textarea viewport.Model
}

// NewModel creates a new pane containing the given model.
func NewLeaf(m tea.Model) Pane {
	borderStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

	return Pane{
		model:    m,
		style:    borderStyle.Align(lipgloss.Center, lipgloss.Center),
		textarea: viewport.New(0, 0),
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

func (m Pane) View() string {
	if m.model != nil {
		style := m.style.Copy().Width(m.width - 2).Height(m.height - 2)
		m.textarea.Width = m.width - 2
		m.textarea.Height = m.height - 2
		m.textarea.SetContent(m.model.View())
		innerView := m.textarea.View()
		if m.name != "" {
			return m.genTop() + "\n" + style.Render(innerView)
		}
		return style.Render(innerView)
	}

	childrenViews := make([]string, len(m.children))
	for i, child := range m.children {
		childrenViews[i] = child.View()
	}

	switch m.direction {
	case DirectionVertical:
		return lipgloss.JoinVertical(lipgloss.Top, childrenViews...)

	case DirectionHorizontal:
		return lipgloss.JoinHorizontal(lipgloss.Left, childrenViews...)

	default:
		panic("unknown direction")
	}
}
