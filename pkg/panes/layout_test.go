package panes_test

import (
	"testing"

	"github.com/evertras/yakdash/pkg/panes"
	"github.com/stretchr/testify/assert"
)

func TestFullPaneTakesFullDimensions(t *testing.T) {
	// Ensure a single pane with no children takes the entire dimensions
	pane, _ := panes.NewLeaf(nil).WithDimensions(200, 100)
	assert.Equal(t, 200, pane.Width(), "Wrong width")
	assert.Equal(t, 100, pane.Height(), "Wrong height")
}

func TestHorizontalPanesSplitHorizontally(t *testing.T) {
	// Ensure a single pane with ten children horizontally splits the width
	pane, _ := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(nil), panes.NewLeaf(nil)).WithDimensions(100, 100)
	for i, child := range pane.Children() {
		assert.Equal(t, 50, child.Width(), "Wrong width for child %d", i)
		assert.Equal(t, 100, child.Height(), "Wrong height for child %d", i)
	}
}

func TestVerticalPanesSplitVertically(t *testing.T) {
	// Ensure a single pane with ten children vertically splits the height
	pane, _ := panes.NewNode(panes.DirectionVertical, panes.NewLeaf(nil), panes.NewLeaf(nil)).WithDimensions(100, 100)
	for i, child := range pane.Children() {
		assert.Equal(t, 100, child.Width(), "Wrong width for child %d", i)
		assert.Equal(t, 50, child.Height(), "Wrong height for child %d", i)
	}
}

func TestNestedPanesSplitWithinParent(t *testing.T) {
	// With a parent pane that contains one leaf child and one node child, split
	// horizontally, which contains ten children, split vertically, ensure the
	// expected dimensions add up to the overall width/height.

	sideChildren := make([]panes.Pane, 10)
	for i := range sideChildren {
		sideChildren[i] = panes.NewLeaf(nil)
	}

	pane, _ := panes.NewNode(panes.DirectionHorizontal,
		panes.NewLeaf(nil),
		panes.NewNode(panes.DirectionVertical, sideChildren...),
	).WithDimensions(100, 100)

	topChildren := pane.Children()
	mainPane := topChildren[0]
	sidePane := topChildren[1]

	assert.Equal(t, 50, mainPane.Width(), "Wrong width for main pane")
	assert.Equal(t, 100, mainPane.Height(), "Wrong height for main pane")

	assert.Equal(t, 50, sidePane.Width(), "Wrong width for side pane")
	assert.Equal(t, 100, sidePane.Height(), "Wrong height for side pane")

	for i, child := range sidePane.Children() {
		assert.Equal(t, 50, child.Width(), "Wrong width for side pane child %d", i)
		assert.Equal(t, 10, child.Height(), "Wrong height for side pane child %d", i)
	}
}

func TestChangingDirectionChangesLayout(t *testing.T) {
	// With a parent pane that contains two leaf children split horizontally,
	// switch to vertical and make sure the layout changes.
	pane, _ := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(nil), panes.NewLeaf(nil)).WithDimensions(100, 100)

	for _, child := range pane.Children() {
		assert.Equal(t, 50, child.Width(), "Wrong width for child")
		assert.Equal(t, 100, child.Height(), "Wrong height for child")
	}

	pane = pane.WithDirection(panes.DirectionVertical)

	for _, child := range pane.Children() {
		assert.Equal(t, 100, child.Width(), "Wrong width for child")
		assert.Equal(t, 50, child.Height(), "Wrong height for child")
	}
}
