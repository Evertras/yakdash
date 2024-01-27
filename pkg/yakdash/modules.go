package yakdash

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/evertras/yakdash/pkg/layout"
	"github.com/evertras/yakdash/pkg/modules/clock"
	"github.com/evertras/yakdash/pkg/modules/command"
	"github.com/evertras/yakdash/pkg/modules/text"
)

func loadModule(l layout.Node) (tea.Model, error) {
	switch l.Module {
	case "clock":
		return clock.New(l.Config)

	case "command":
		return command.New(l.Config)

	default:
		return text.New("Unknown module: " + l.Module), nil
	}
}
