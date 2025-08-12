package palette

import "github.com/charmbracelet/lipgloss"

// Color represents a single color with its name and styling information.
type Color struct {
	// Name is the name of the color.
	Name string

	// Hex is the hex code of the color.
	Hex string

	// Style is the lipgloss style for the color.
	Style lipgloss.Style
}

// ColorDefinition represents a color name and hex value pair.
type ColorDefinition struct {
	// Name is the name of the color.
	Name string

	// Hex is the hex code of the color.
	Hex string
}

// NewColor creates a new Color instance.
func NewColor(name, hex string) *Color {
	return &Color{
		Name:  name,
		Hex:   hex,
		Style: renderer.NewStyle().Foreground(lipgloss.Color(hex)),
	}
}
