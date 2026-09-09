//nolint:goconst
package palette

// CreatePoimandresPalette creates the Poimandres color palette.
func CreatePoimandresPalette() *Palette {
	palette := NewPalette("Poimandres", "Poimandres", "minimal", "modern", "dark", "frameless")

	// From https://github.com/olivercederborg/poimandres.nvim
	colors := []ColorDefinition{
		{"yellow", "#fffac2"},
		{"teal1", "#5de4c7"},
		{"teal2", "#5fb3a1"},
		{"teal3", "#42675a"},
		{"blue1", "#89ddff"},
		{"blue2", "#add7ff"},
		{"blue3", "#91b4d5"},
		{"blue4", "#7390aa"},
		{"pink1", "#fae4fc"},
		{"pink2", "#fcc5e9"},
		{"pink3", "#d0679d"},
		{"blueGray1", "#a6accd"},
		{"blueGray2", "#767c9d"},
		{"blueGray3", "#506477"},
		{"background1", "#303340"},
		{"background2", "#1b1e28"},
		{"background3", "#171922"},
		{"text", "#e4f0fb"},
		{"white", "#ffffff"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}
