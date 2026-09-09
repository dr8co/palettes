//nolint:goconst
package palette

// OneDark Pro
func createOneDarkPalette() *Palette {
	palette := NewPalette("OneDark Pro", "OneDark", "OneDark Pro", "One", "dark", "Atom")

	// From https://github.com/Binaryify/OneDark-Pro
	colors := []ColorDefinition{
		{"deepRed", "#be5046"},
		{"error", "#f44747"},
		{"coral", "#e06c75"},
		{"purple", "#c678dd"},
		{"malibu", "#61afef"},
		{"fountainBlue", "#56b6c2"},
		{"green", "#98c379"},
		{"whiskey", "#d19a66"},
		{"chalky", "#e5c07b"},
		{"invalid", "#ffffff"},
		{"lightWhite", "#abb2bf"},
		{"lightDark", "#7f848e"},
		{"dark", "#5c6370"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}

// OneLight
func createOneLightPalette() *Palette {
	palette := NewPalette("OneLight", "OneLight", "One", "light", "Atom")

	// From https://github.com/olimorris/onedarkpro.nvim
	colors := []ColorDefinition{
		{"red", "#e05661"},
		{"orange", "#ee9025"},
		{"yellow", "#eea825"},
		{"highlight", "#e2be7d"},
		{"purple", "#9a77cf"},
		{"blue", "#118dc3"},
		{"cyan", "#56b6c2"},
		{"green", "#1da912"},
		{"white", "#fafafa"},
		{"gray", "#bebebe"},
		{"comment", "#9b9fa6"},
		{"black", "#6a6a6a"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}

// OneDark Pro Vivid
func createOneDarkVividPalette() *Palette {
	palette := NewPalette("OneDark Pro Vivid", "OneDark", "OneDark Pro", "One", "dark", "Atom", "Vivid")

	// From https://github.com/Binaryify/OneDark-Pro
	colors := []ColorDefinition{
		{"deepRed", "#be5046"},
		{"error", "#f44747"},
		{"coral", "#ef596f"},
		{"purple", "#d55fde"},
		{"malibu", "#61afef"},
		{"fountainBlue", "#2bbac5"},
		{"green", "#89ca78"},
		{"whiskey", "#d19a66"},
		{"chalky", "#e5c07b"},
		{"invalid", "#ffffff"},
		{"lightWhite", "#abb2bf"},
		{"lightDark", "#7f848e"},
		{"dark", "#5c6370"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}

// Vaporwave
func createVaporwavePalette() *Palette {
	palette := NewPalette("Vaporwave", "OneDark", "One", "dark", "Atom", "Vaporwave")

	// From https://github.com/olimorris/onedarkpro.nvim
	colors := []ColorDefinition{
		{"red", "#e16765"},
		{"orange", "#eaa041"},
		{"yellow", "#eae852"},
		{"highlight", "#e2be7d"},
		{"purple", "#c678dd"},
		{"blue", "#25abe4"},
		{"cyan", "#46a3af"},
		{"green", "#75be78"},
		{"white", "#b4b7cf"},
		{"comment", "#7679A7"},
		{"gray", "#585b89"},
		{"black", "#222435"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}

// CreateAtomPalettes creates a set of Atom palettes.
func CreateAtomPalettes() []*Palette {
	return []*Palette{
		createOneDarkPalette(),
		createOneLightPalette(),
		createOneDarkVividPalette(),
		createVaporwavePalette(),
	}
}
