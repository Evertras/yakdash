package panes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullPaneTakesFullDimensions(t *testing.T) {
	// Ensure a single pane with no children takes the entire dimensions
	pane := NewLeaf(nil).WithDimensions(200, 100)
	assert.Equal(t, 200, pane.width, "Wrong width")
	assert.Equal(t, 100, pane.height, "Wrong height")
}

func TestHorizontalPanesSplitHorizontally(t *testing.T) {
	// Ensure a single pane with ten children horizontally splits the width
	pane := NewNode(DirectionHorizontal, NewLeaf(nil), NewLeaf(nil)).WithDimensions(100, 100)
	for i, child := range pane.children {
		assert.Equal(t, 50, child.width, "Wrong width for child %d", i)
		assert.Equal(t, 100, child.height, "Wrong height for child %d", i)
	}
}

func TestVerticalPanesSplitVertically(t *testing.T) {
	// Ensure a single pane with ten children vertically splits the height
	pane := NewNode(DirectionVertical, NewLeaf(nil), NewLeaf(nil)).WithDimensions(100, 100)
	for i, child := range pane.children {
		assert.Equal(t, 100, child.width, "Wrong width for child %d", i)
		assert.Equal(t, 50, child.height, "Wrong height for child %d", i)
	}
}

func TestNestedPanesSplitWithinParent(t *testing.T) {
	// With a parent pane that contains one leaf child and one node child, split
	// horizontally, which contains ten children, split vertically, ensure the
	// expected dimensions add up to the overall width/height.

	children := make([]Pane, 10)
	for i := range children {
		children[i] = NewLeaf(nil)
	}

	pane := NewNode(DirectionHorizontal,
		NewLeaf(nil),
		NewNode(DirectionVertical, children...),
	).WithDimensions(100, 100)

	mainPane := pane.children[0]
	sidePane := pane.children[1]

	assert.Equal(t, 50, mainPane.width, "Wrong width for main pane")
	assert.Equal(t, 100, mainPane.height, "Wrong height for main pane")

	assert.Equal(t, 50, sidePane.width, "Wrong width for side pane")
	assert.Equal(t, 100, sidePane.height, "Wrong height for side pane")

	for i, child := range sidePane.children {
		assert.Equal(t, 50, child.width, "Wrong width for side pane child %d", i)
		assert.Equal(t, 10, child.height, "Wrong height for side pane child %d", i)
	}
}

func TestChangingDirectionChangesLayout(t *testing.T) {
	// With a parent pane that contains two leaf children split horizontally,
	// switch to vertical and make sure the layout changes.
	pane := NewNode(DirectionHorizontal, NewLeaf(nil), NewLeaf(nil)).WithDimensions(100, 100)

	assert.Equal(t, 50, pane.children[0].width, "Wrong width for child 0")

	pane = pane.WithDirection(DirectionVertical)

	assert.Equal(t, 100, pane.children[0].width, "Wrong width for child 0")
}
