package palette

import "github.com/charmbracelet/lipgloss"

type Color struct {
	Name  string
	Hex   string
	Style lipgloss.Style
}

type ColorDefinition struct {
	Name string
	Hex  string
}

func NewColor(name, hex string) *Color {
	return &Color{
		Name:  name,
		Hex:   hex,
		Style: renderer.NewStyle().Foreground(lipgloss.Color(hex)),
	}
}
