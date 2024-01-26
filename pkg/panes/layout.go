package panes

func (m Pane) recalculateDimensions() Pane {
	numChildren := len(m.children)
	if numChildren == 0 {
		return m
	}
	width := m.width
	height := m.height
	var overflow int

	if m.direction == DirectionVertical {
		height = height / numChildren
		overflow = height % numChildren
	} else {
		width = width / numChildren
		overflow = width % numChildren
	}

	for i := range m.children {
		// Shadow to allow adjustment below
		height := height
		width := width

		if i == numChildren-1 {
			if m.direction == DirectionVertical {
				height += overflow
			} else {
				width += overflow
			}
		}

		m.children[i] = m.children[i].WithDimensions(width, height)
	}

	return m
}
