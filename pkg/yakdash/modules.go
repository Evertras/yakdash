package yakdash

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/evertras/yakdash/pkg/layout"
	"github.com/evertras/yakdash/pkg/modules/clock"
	"github.com/evertras/yakdash/pkg/modules/text"
)

func loadModule(l layout.Node) tea.Model {
	switch l.Module {
	case "clock":
		return clock.New()

	default:
		return text.New("Unknown module: " + l.Module)
	}
}
