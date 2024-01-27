package panes

func (m Pane) WithDimensions(width, height int) Pane {
	m.width = width
	m.height = height

	m = m.recalculateDimensions()

	return m
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
