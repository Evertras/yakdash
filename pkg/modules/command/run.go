package command

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/evertras/yakdash/pkg/id"
	"github.com/evertras/yakdash/pkg/sources/capturedexec"
)

type commandResult struct {
	id     id.ID
	stdout string
	stderr string
	err    error
}

func (m model) runCommand() tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), m.config.interval)
		defer cancel()
		cmd := capturedexec.New(ctx, "bash", "-c", m.config.Bash)

		err := cmd.Run()

		return commandResult{
			id:     m.id,
			stdout: cmd.Stdout(),
			stderr: cmd.Stderr(),
			err:    err,
		}
	}
}
