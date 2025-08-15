package palette

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Color represents a single color with its definition and styling information.
type Color struct {
	// Def is the color definition (name and hex).
	Def ColorDefinition

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
		Def:   ColorDefinition{Name: name, Hex: hex},
		Style: renderer.NewStyle().Foreground(lipgloss.Color(hex)),
	}
}

// String returns a string representation of a [ColorDefinition].
func (cd *ColorDefinition) String() string {
	if cd.Name == "" {
		return cd.Hex
	}
	return fmt.Sprintf("%s (%s)", cd.Name, cd.Hex)
}
