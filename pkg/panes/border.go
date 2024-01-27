package panes

import "strings"

func (m *Pane) genTop() string {
	builder := strings.Builder{}
	targetFillWidth := m.width - 2
	targetNameWidth := len(m.name) + 2

	if targetFillWidth <= 0 {
		targetFillWidth = 0
	}

	if targetNameWidth > targetFillWidth {
		targetNameWidth = targetFillWidth
	}

	targetFillWidth -= targetNameWidth

	builder.WriteRune('╭')
	builder.WriteRune('╮')
	builder.WriteString(m.name)
	builder.WriteRune('╭')
	builder.WriteString(strings.Repeat("─", targetFillWidth))
	builder.WriteRune('╮')

	return builder.String()
}
