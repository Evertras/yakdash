package panes_test

import (
	"testing"

	"github.com/evertras/yakdash/pkg/panes"
	"github.com/stretchr/testify/assert"
)

func TestAlignmentVertical(t *testing.T) {
	testCases := []struct {
		raw      string
		expected panes.AlignmentVertical
	}{
		{"", panes.AlignmentVerticalCenter},
		{"top", panes.AlignmentVerticalTop},
		{"center", panes.AlignmentVerticalCenter},
		{"bottom", panes.AlignmentVerticalBottom},
	}

	for _, tc := range testCases {
		t.Run(tc.raw, func(t *testing.T) {
			actual, err := panes.ToAlignmentVertical(tc.raw)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestAlignmentVerticalError(t *testing.T) {
	_, err := panes.ToAlignmentVertical("asdf")

	assert.Error(t, err)
}

func TestAlignmentHorizontal(t *testing.T) {
	testCases := []struct {
		raw      string
		expected panes.AlignmentHorizontal
	}{
		{"", panes.AlignmentHorizontalCenter},
		{"left", panes.AlignmentHorizontalLeft},
		{"center", panes.AlignmentHorizontalCenter},
		{"right", panes.AlignmentHorizontalRight},
	}

	for _, tc := range testCases {
		t.Run(tc.raw, func(t *testing.T) {
			actual, err := panes.ToAlignmentHorizontal(tc.raw)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestAlignmentHorizontalError(t *testing.T) {
	_, err := panes.ToAlignmentHorizontal("asdf")

	assert.Error(t, err)
}

func TestWithAlignment(t *testing.T) {
	testCases := []struct {
		name       string
		vertical   panes.AlignmentVertical
		horizontal panes.AlignmentHorizontal

		expectedViewOutput string
	}{
		{
			name:       "centered",
			vertical:   panes.AlignmentVerticalCenter,
			horizontal: panes.AlignmentHorizontalCenter,
			expectedViewOutput: `
╭────────────╮
│            │
│    foo     │
│            │
╰────────────╯`,
		},
		{
			name:       "top left",
			vertical:   panes.AlignmentVerticalTop,
			horizontal: panes.AlignmentHorizontalLeft,
			expectedViewOutput: `
╭────────────╮
│foo         │
│            │
│            │
╰────────────╯`,
		},
		{
			name:       "bottom right",
			vertical:   panes.AlignmentVerticalBottom,
			horizontal: panes.AlignmentHorizontalRight,
			expectedViewOutput: `
╭────────────╮
│            │
│            │
│         foo│
╰────────────╯`,
		},
		{
			name:       "top right",
			vertical:   panes.AlignmentVerticalTop,
			horizontal: panes.AlignmentHorizontalRight,
			expectedViewOutput: `
╭────────────╮
│         foo│
│            │
│            │
╰────────────╯`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, _ := panes.NewLeaf(newDummyModel("foo", nil, nil)).WithDimensions(14, 5)

			p = p.WithAlignment(tc.vertical, tc.horizontal)

			// Trim the leading \n in all the test cases, have it
			// that way for readability
			assert.Equal(t, tc.expectedViewOutput[1:], p.View())
		})
	}
}
