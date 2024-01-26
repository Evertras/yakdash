package yakdash

import (
	"github.com/evertras/yakdash/pkg/layout"
	"github.com/evertras/yakdash/pkg/modules/clock"
	"github.com/evertras/yakdash/pkg/panes"
)

func New(l layout.Root) model {
	var p []panes.Pane

	var makePane func(node layout.Node) panes.Pane
	makePane = func(node layout.Node) panes.Pane {
		if len(node.Children) > 0 {
			children := make([]panes.Pane, len(node.Children))
			for i, child := range node.Children {
				children[i] = makePane(child)
			}

			direction := panes.DirectionVertical
			if node.Stack == "horizontal" {
				direction = panes.DirectionHorizontal
			}

			return panes.NewNode(direction, children...)
		}

		// TODO: load module
		return panes.NewLeaf(clock.New())
	}

	for _, node := range l.Screens {
		p = append(p, makePane(node))
	}

	return model{
		layout:   l,
		rootPane: makePane(l.Screens[0]),
	}
}
