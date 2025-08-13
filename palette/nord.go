package palette

import "strings"

// NordColors contains the color definitions for the Nord theme.
// Check https://www.nordtheme.com/docs/colors-and-palettes for more information.
var NordColors = map[string][]ColorDefinition{
	"polar night": {
		{"nord0", "#2e3440"},
		{"nord1", "#3b4252"},
		{"nord2", "#434c5e"},
		{"nord3", "#4c566a"},
	},
	"snow storm": {
		{"nord4", "#d8dee9"},
		{"nord5", "#e5e9f0"},
		{"nord6", "#eceff4"},
	},
	"frost": {
		{"nord7", "#8fbcbb"},
		{"nord8", "#88c0d0"},
		{"nord9", "#81a1c1"},
		{"nord10", "#5e81ac"},
	},
	"aurora": {
		{"nord11", "#bf616a"},
		{"nord12", "#d08770"},
		{"nord13", "#ebcb8b"},
		{"nord14", "#a3be8c"},
		{"nord15", "#b48ead"},
	},
}

// CreateNordPalette creates a Nord palette variant with appropriate families.
func CreateNordPalette(variant string) *Palette {
	variant = strings.ToLower(strings.TrimSpace(variant))
	colors, exists := NordColors[variant]
	if !exists {
		return nil
	}

	name := "Nord " + variant

	palette := NewPalette(name, "Nord", variant)

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}
	return palette
}
