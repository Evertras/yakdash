package panes_test

import tea "github.com/charmbracelet/bubbletea"

type dummyModel struct {
	text string

	initCmd        func() tea.Msg
	updateCallback func(tea.Msg)
}

func newDummyModel(text string, initCmd func() tea.Msg, updateCallback func(tea.Msg)) dummyModel {
	return dummyModel{
		text: text,

		initCmd:        initCmd,
		updateCallback: updateCallback,
	}
}

func (m dummyModel) Init() tea.Cmd {
	m.initCmd()

	return m.initCmd
}

func (m dummyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.updateCallback(msg)

	return m, nil
}

func (m dummyModel) View() string {
	return m.text
}
