package panes_test

import tea "github.com/charmbracelet/bubbletea"

type dummyModel struct {
	text           string
	updateCallback func(tea.Msg)
}

func newDummyModel(text string, updateCallback func(tea.Msg)) dummyModel {
	return dummyModel{
		text:           text,
		updateCallback: updateCallback,
	}
}

func (m dummyModel) Init() tea.Cmd {
	return nil
}

func (m dummyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.updateCallback(msg)

	return m, nil
}

func (m dummyModel) View() string {
	return m.text
}
