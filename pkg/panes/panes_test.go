package panes_test

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"

	"github.com/evertras/yakdash/pkg/panes"
)

func TestUpdateLeafNodeUpdatesInnerModel(t *testing.T) {
	// Given a leaf node with a simple model,
	// the model should receive the update message
	sawUpdate := false
	dummy := newDummyModel("testing", func(tea.Msg) { sawUpdate = true })
	pane := panes.NewLeaf(dummy)

	pane.Update(7)

	assert.True(t, sawUpdate, "Model should have seen update")
}

func TestUpdateParentNodeSendsUpdateToChildren(t *testing.T) {
	// Given a parent node with nested children,
	// the model should receive the update message
	sawUpdate := false
	dummy := newDummyModel("testing", func(tea.Msg) { sawUpdate = true })
	parent := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(dummy))

	parent.Update(7)

	assert.True(t, sawUpdate, "Model should have seen update")
}

func TestViewLeafNodeShowsInnerModel(t *testing.T) {
	// Given a leaf node with a simple model,
	// the view should return the inner model's view
	dummy := newDummyModel("testing", nil)
	pane := panes.NewLeaf(dummy).WithDimensions(100, 100)

	assert.Contains(t, pane.View(), "testing", "View should return inner model's view")
}

func TestViewParentNodeShowsInnerModelsOfChildren(t *testing.T) {
	// Given a leaf node with a simple model,
	// the view should return the inner model's view
	dummyLeft := newDummyModel("left", nil)
	dummyRight := newDummyModel("right", nil)
	pane := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(dummyLeft), panes.NewLeaf(dummyRight)).WithDimensions(100, 1)

	assert.Contains(t, pane.View(), "left", "View should return left child")
	assert.Contains(t, pane.View(), "right", "View should return right child")
}

func TestViewHorizontalGoesLeftToRight(t *testing.T) {
	dummyLeft := newDummyModel("left", nil)
	dummyRight := newDummyModel("right", nil)
	pane := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(dummyLeft), panes.NewLeaf(dummyRight)).WithDimensions(100, 1)

	indexLeft := strings.Index(pane.View(), "left")
	indexRight := strings.Index(pane.View(), "right")

	assert.True(t, indexLeft < indexRight, "Left child should be before right child")
}

func TestViewVerticalGoesTopToBottom(t *testing.T) {
	dummyTop := newDummyModel("top", nil)
	dummyBottom := newDummyModel("bottom", nil)
	pane := panes.NewNode(panes.DirectionVertical, panes.NewLeaf(dummyTop), panes.NewLeaf(dummyBottom)).WithDimensions(10, 100)

	indexTop := strings.Index(pane.View(), "top")
	indexBottom := strings.Index(pane.View(), "bottom")

	assert.True(t, indexTop < indexBottom, "Top child should be before bottom child")
}
