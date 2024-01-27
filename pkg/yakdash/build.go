package yakdash

import (
	"fmt"

	"github.com/evertras/yakdash/pkg/layout"
	"github.com/evertras/yakdash/pkg/modules/text"
	"github.com/evertras/yakdash/pkg/panes"
)

func New(l layout.Root) model {
	var allPanes []panes.Pane

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

		module, err := loadModule(node)

		if err != nil {
			module = text.New(fmt.Sprintf("Error loading module %q: %s", node.Module, err.Error()))
		}

		return panes.NewLeaf(module)
	}

	for _, node := range l.Screens {
		allPanes = append(allPanes, makePane(node))
	}

	return model{
		layout:   l,
		rootPane: allPanes[0],
	}
}
