package panes_test

import (
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
