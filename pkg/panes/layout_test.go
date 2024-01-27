package panes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullPaneTakesFullDimensions(t *testing.T) {
	// Ensure a single pane with no children takes the entire dimensions
	pane := NewLeaf(nil).WithDimensions(100, 100)
	pane.recalculateDimensions()
	assert.Equal(t, 100, pane.width, "Wrong width")
	assert.Equal(t, 100, pane.height, "Wrong height")
}

func TestHorizontalPanesSplitHorizontally(t *testing.T) {
	// Ensure a single pane with ten children horizontally splits the width
	pane := NewNode(DirectionHorizontal, NewLeaf(nil), NewLeaf(nil)).WithDimensions(100, 100)
	pane.recalculateDimensions()
	for i, child := range pane.children {
		assert.Equal(t, 50, child.width, "Wrong width for child %d", i)
		assert.Equal(t, 100, child.height, "Wrong height for child %d", i)
	}
}

func TestVerticalPanesSplitVertically(t *testing.T) {
	// Ensure a single pane with ten children vertically splits the height
	pane := NewNode(DirectionVertical, NewLeaf(nil), NewLeaf(nil)).WithDimensions(100, 100)
	pane.recalculateDimensions()
	for i, child := range pane.children {
		assert.Equal(t, 100, child.width, "Wrong width for child %d", i)
		assert.Equal(t, 50, child.height, "Wrong height for child %d", i)
	}
}
