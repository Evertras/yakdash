package panes

func (m Pane) recalculateDimensions() Pane {
	numChildren := len(m.children)
	if numChildren == 0 {
		return m
	}
	width := m.width
	height := m.height

	if m.direction == DirectionVertical {
		height = height / numChildren
	} else {
		width = width / numChildren
	}

	for i := range m.children {
		// Shadow to allow adjustment below
		height := height
		width := width

		m.children[i] = m.children[i].WithDimensions(width, height)
	}

	return m
}
