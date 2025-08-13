package palette

// CreateEldritchPalette creates the Eldritch color palette.
func CreateEldritchPalette() *Palette {
	palette := NewPalette("Eldritch", "Eldritch", "dark", "Lovecraft")

	// From https://github.com/eldritch-theme/eldritch
	colors := []ColorDefinition{
		{"Gold of Yuggoth", "#f1fc79"},
		{"R'lyeh' Red", "#f16c75"},
		{"Lovecraft Purple", "#a48cf2"},
		{"Pustule Pink", "#f265b5"},
		{"Dreaming Orange", "#f7c67f"},
		{"Great Old One Green", "#37f499"},
		{"Watery Tomb Blue", "#04d1f9"},
		{"The Old One Purple", "#7081d0"},
		{"Lighthouse White", "#ebfafa"},
		{"Shallow Depths Grey", "#323449"},
		{"Sunken Depths Grey", "#212337"},
	}

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}

	return palette
}
