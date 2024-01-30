package panes

import (
	"bufio"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Pane) View() string {
	if m.model != nil {
		style := m.style.Copy().Width(m.width - 2).Height(m.height - 2)
		innerView := m.crop(m.model.View())
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

func (m Pane) crop(str string) string {
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)

	cropped := strings.Builder{}

	maxCrop := m.width - 2
	first := true

	for i := 0; i < m.height-2; i++ {
		if !scanner.Scan() {
			break
		}
		if !first {
			cropped.WriteRune('\n')
		}
		first = false
		line := scanner.Text()

		crop := maxCrop
		if len(line) < maxCrop {
			cropped.WriteString(line)
		} else {
			cropped.WriteString(scanner.Text()[:crop])
		}
	}

	return cropped.String()
}
