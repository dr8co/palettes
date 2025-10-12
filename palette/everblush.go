package palette

// CreateEverblushPalette creates the Everblush color palette.
func CreateEverblushPalette() *Palette {
	palette := NewPalette("Everblush", "Everblush", "dark", "pastel")

	// From https://github.com/Everblush
	colors := []ColorDefinition{
		{"red", "#e57474"},
		{"green", "#8ccf7e"},
		{"yellow", "#e5c76b"},
		{"blue", "#67b0e8"},
		{"magenta", "#c47fd5"},
		{"cyan", "#6cbfbf"},
		{"white", "#dadada"},
		{"light gray", "#b3b9b8"},
		{"lighter background", "#232a2d"},
		{"background", "#141b1e"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}
