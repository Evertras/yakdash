package panes

import (
	"bufio"
	"strings"
)

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
