package command

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/evertras/yakdash/pkg/id"
	"github.com/mitchellh/mapstructure"
)

type model struct {
	id     id.ID
	config config

	lastOutput string
}

func New(cfg interface{}) (model, error) {
	var config config
	err := mapstructure.Decode(cfg, &config)

	if err != nil {
		return model{}, fmt.Errorf("failed to decode config: %w", err)
	}

	config, err = config.setDefaultsAndParse()

	if err != nil {
		return model{}, fmt.Errorf("invalid config: %w", err)
	}

	m := model{
		id:     id.New(),
		config: config,
	}

	return m, nil
}

type tickMsg struct {
	id id.ID
}

func (m model) doTick() tea.Cmd {
	// Tick on a regular interval so that all commands
	// from different panes can try to update at the
	// same time, if they use the same interval.
	// The command will timeout before the next tick.
	return tea.Batch(
		tea.Tick(m.config.interval, func(time.Time) tea.Msg {
			return tickMsg{m.id}
		}),
		m.runCommand(),
	)
}

func (m model) Init() tea.Cmd {
	return m.doTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tickMsg:
		if msg.id != m.id {
			break
		}

		cmds = append(cmds, m.doTick())

	case commandResult:
		if msg.id != m.id {
			break
		}

		if msg.err != nil {
			m.lastOutput = fmt.Sprintf("Error: %s", msg.err)
			break
		}

		m.lastOutput = msg.stdout

	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.lastOutput
}
