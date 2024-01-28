package panes

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type AlignmentVertical string
type AlignmentHorizontal string

const (
	AlignmentVerticalTop    AlignmentVertical = "top"
	AlignmentVerticalCenter AlignmentVertical = "center"
	AlignmentVerticalBottom AlignmentVertical = "bottom"

	AlignmentHorizontalLeft   AlignmentHorizontal = "left"
	AlignmentHorizontalCenter AlignmentHorizontal = "center"
	AlignmentHorizontalRight  AlignmentHorizontal = "right"
)

func ToAlignmentVertical(raw string) (AlignmentVertical, error) {
	if raw == "" {
		return AlignmentVerticalCenter, nil
	}

	switch AlignmentVertical(raw) {
	case AlignmentVerticalTop:
		return AlignmentVerticalTop, nil
	case AlignmentVerticalCenter:
		return AlignmentVerticalCenter, nil
	case AlignmentVerticalBottom:
		return AlignmentVerticalBottom, nil
	default:
		return "", fmt.Errorf("invalid vertical alignment: %s", raw)
	}
}

func ToAlignmentHorizontal(raw string) (AlignmentHorizontal, error) {
	if raw == "" {
		return AlignmentHorizontalCenter, nil
	}

	switch AlignmentHorizontal(raw) {
	case AlignmentHorizontalLeft:
		return AlignmentHorizontalLeft, nil
	case AlignmentHorizontalCenter:
		return AlignmentHorizontalCenter, nil
	case AlignmentHorizontalRight:
		return AlignmentHorizontalRight, nil
	default:
		return "", fmt.Errorf("invalid horizontal alignment: %s", raw)
	}
}

func (m Pane) WithAlignment(vertical AlignmentVertical, horizontal AlignmentHorizontal) Pane {
	var (
		lipglossV lipgloss.Position
		lipglossH lipgloss.Position
	)

	switch vertical {
	case AlignmentVerticalTop:
		lipglossV = lipgloss.Top
	case AlignmentVerticalCenter:
		lipglossV = lipgloss.Center
	case AlignmentVerticalBottom:
		lipglossV = lipgloss.Bottom
	}

	switch horizontal {
	case AlignmentHorizontalLeft:
		lipglossH = lipgloss.Left
	case AlignmentHorizontalCenter:
		lipglossH = lipgloss.Center
	case AlignmentHorizontalRight:
		lipglossH = lipgloss.Right
	}

	m.style = m.style.Align(lipglossH, lipglossV)

	return m
}
