package yakdash

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/evertras/yakdash/pkg/layout"
)

type model struct {
	layout layout.Root
}

func New(layout layout.Root) model {
	return model{
		layout: layout,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	builder := strings.Builder{}
	ws := func(s string) {
		builder.WriteString(s)
	}
	endline := func() {
		builder.WriteString("\n")
	}

	var printNode func(node layout.Node, depth int)
	printNode = func(node layout.Node, depth int) {
		indented := func() {
			for i := 0; i < depth; i++ {
				ws(" ")
			}
		}

		indented()
		ws("- Node: ")
		ws(node.Name)
		endline()

		if len(node.Children) > 0 {
			for _, node := range node.Children {
				printNode(node, depth+2)
			}
		} else {
			indented()
			ws("  Module: ")
			ws(node.Module)
			endline()
		}
	}

	for _, node := range m.layout.Screens {
		printNode(node, 0)
	}

	return builder.String()
}
