package palette

// CreateDraculaPalette creates the Dracula color palette.
func CreateDraculaPalette() *Palette {
	palette := NewPalette("Dracula", "Dracula")

	// From https://github.com/dracula/dracula-theme
	colors := []ColorDefinition{
		{"background", "#282a36"},
		{"current line", "#44475a"},
		{"selection", "#44475a"},
		{"foreground", "#f8f8f2"},
		{"comment", "#6272a4"},
		{"cyan", "#8be9fd"},
		{"green", "#50fa7b"},
		{"orange", "#ffb86c"},
		{"pink", "#ff79c6"},
		{"purple", "#bd93f9"},
		{"red", "#ff5555"},
		{"yellow", "#f1fa8c"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}
