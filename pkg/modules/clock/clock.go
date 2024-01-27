package clock

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/yakdash/pkg/id"
	"github.com/mitchellh/mapstructure"
)

type model struct {
	id       id.ID
	config   config
	text     string
	location *time.Location
}

func New(cfg interface{}) (model, error) {
	var config config
	err := mapstructure.Decode(cfg, &config)

	if err != nil {
		return model{}, fmt.Errorf("failed to decode config: %w", err)
	}

	m := model{
		id:     id.New(),
		config: config.fillDefaults(),
	}

	m.location, err = time.LoadLocation(m.config.Timezone)

	if err != nil {
		return model{}, fmt.Errorf("failed to load timezone %q: %w", m.config.Timezone, err)
	}

	m = m.updateText()

	return m, nil
}

type tickMsg struct {
	id id.ID
}

func doTick(id id.ID) tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{id}
	})
}

func (m model) updateText() model {
	// Get the time based on the model's timezone
	m.text = time.Now().In(m.location).Format(m.config.Format)

	return m
}

func (m model) Init() tea.Cmd {
	return doTick(m.id)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if msg.id != m.id {
			break
		}
		m = m.updateText()
		return m, doTick(m.id)
	}

	return m, nil
}

func (m model) View() string {
	return m.text
}
