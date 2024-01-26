package panes

import (
	"math/rand"

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

	width  int
	height int
}

// NewModel creates a new pane containing the given model.
func NewLeaf(m tea.Model) Pane {
	return Pane{
		model: m,
	}
}

// NewNode creates a new pane containing the given children.
func NewNode(direction Direction, children ...Pane) Pane {
	return Pane{
		direction: direction,
		children:  children,
	}
}

func (m Pane) WithDimensions(width, height int) Pane {
	m.width = width
	m.height = height

	m = m.recalculateDimensions()

	return m
}

func (m Pane) WithDirection(direction Direction) Pane {
	m.direction = direction

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
	if m.model != nil {
		style := lipgloss.NewStyle().
			Width(m.width).Height(m.height).
			Background(lipgloss.Color(randomColorHex())).
			Align(lipgloss.Center, lipgloss.Center)
		return style.Render(m.model.View())
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
