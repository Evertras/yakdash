package panes_test

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"

	"github.com/evertras/yakdash/pkg/panes"
)

func TestInitIncludesInnerModel(t *testing.T) {
	// Given a leaf node with a simple model,
	// the model should be initialized
	calledInitLeft := false
	calledInitRight := false
	dummyLeft := newDummyModel("testing", func() tea.Msg { calledInitLeft = true; return "left" }, nil)
	dummyRight := newDummyModel("testing", func() tea.Msg { calledInitRight = true; return "right" }, nil)

	pane := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(dummyLeft), panes.NewLeaf(dummyRight))

	cmd := pane.Init()

	assert.True(t, calledInitLeft, "Left model should have been initialized")
	assert.True(t, calledInitRight, "Right model should have been initialized")

	msgs := cmd()

	switch msgs := msgs.(type) {
	case tea.BatchMsg:
		assert.Equal(t, 2, len(msgs), "Should have gotten two messages")
	default:
		assert.Fail(t, "Should have gotten a slice of messages")
	}
}

func TestUpdateLeafNodeUpdatesInnerModel(t *testing.T) {
	// Given a leaf node with a simple model,
	// the model should receive the update message
	sawUpdate := false
	dummy := newDummyModel("testing", nil, func(tea.Msg) { sawUpdate = true })
	pane := panes.NewLeaf(dummy)

	pane.Update(7)

	assert.True(t, sawUpdate, "Model should have seen update")
}

func TestUpdateParentNodeSendsUpdateToChildren(t *testing.T) {
	// Given a parent node with nested children,
	// the model should receive the update message
	sawUpdate := false
	dummy := newDummyModel("testing", nil, func(tea.Msg) { sawUpdate = true })
	parent := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(dummy))

	parent.Update(7)

	assert.True(t, sawUpdate, "Model should have seen update")
}

func TestViewLeafNodeShowsInnerModel(t *testing.T) {
	// Given a leaf node with a simple model,
	// the view should return the inner model's view
	dummy := newDummyModel("testing", nil, nil)
	pane := panes.NewLeaf(dummy).WithDimensions(100, 100)

	assert.Contains(t, pane.View(), "testing", "View should return inner model's view")
}

func TestViewParentNodeShowsInnerModelsOfChildren(t *testing.T) {
	// Given a leaf node with a simple model,
	// the view should return the inner model's view
	dummyLeft := newDummyModel("left", nil, nil)
	dummyRight := newDummyModel("right", nil, nil)
	pane := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(dummyLeft), panes.NewLeaf(dummyRight)).WithDimensions(100, 1)

	assert.Contains(t, pane.View(), "left", "View should return left child")
	assert.Contains(t, pane.View(), "right", "View should return right child")
}

func TestViewHorizontalGoesLeftToRight(t *testing.T) {
	dummyLeft := newDummyModel("left", nil, nil)
	dummyRight := newDummyModel("right", nil, nil)
	pane := panes.NewNode(panes.DirectionHorizontal, panes.NewLeaf(dummyLeft), panes.NewLeaf(dummyRight)).WithDimensions(100, 1)

	indexLeft := strings.Index(pane.View(), "left")
	indexRight := strings.Index(pane.View(), "right")

	assert.True(t, indexLeft < indexRight, "Left child should be before right child")
}

func TestViewVerticalGoesTopToBottom(t *testing.T) {
	dummyTop := newDummyModel("top", nil, nil)
	dummyBottom := newDummyModel("bottom", nil, nil)
	pane := panes.NewNode(panes.DirectionVertical, panes.NewLeaf(dummyTop), panes.NewLeaf(dummyBottom)).WithDimensions(10, 100)

	indexTop := strings.Index(pane.View(), "top")
	indexBottom := strings.Index(pane.View(), "bottom")

	assert.True(t, indexTop < indexBottom, "Top child should be before bottom child")
}
