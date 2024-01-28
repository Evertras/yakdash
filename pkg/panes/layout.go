package panes

import tea "github.com/charmbracelet/bubbletea"

type ViewableSize struct {
	Width  int
	Height int
}

func (m Pane) WithDimensions(width, height int) (Pane, tea.Cmd) {
	m.width = width
	m.height = height

	return m.recalculateDimensions()
}

func (m Pane) WithDirection(direction Direction) Pane {
	m.direction = direction

	m.recalculateDimensions()

	return m
}

func (m Pane) Width() int {
	return m.width
}

func (m Pane) Height() int {
	return m.height
}

func (m Pane) Children() []Pane {
	return m.children
}

func (m Pane) recalculateDimensions() (Pane, tea.Cmd) {
	numChildren := len(m.children)
	if numChildren == 0 {
		var cmd tea.Cmd = nil
		if m.model != nil {
			m.model, cmd = m.model.Update(ViewableSize{
				Width:  m.width - 2,
				Height: m.height - 2,
			})
		}
		return m, cmd
	}
	width := m.width
	height := m.height

	if m.direction == DirectionVertical {
		height = height / numChildren
	} else {
		width = width / numChildren
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	for i := range m.children {
		m.children[i], cmd = m.children[i].WithDimensions(width, height)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
