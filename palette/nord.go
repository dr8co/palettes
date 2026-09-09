package palette

// Contains the color definitions for the Nord theme.
// Check https://www.nordtheme.com/docs/colors-and-palettes for more information.
var nordColors = map[string][]ColorDefinition{
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

// CreateNordPalettes creates a set of Nord palettes.
func CreateNordPalettes() []*Palette {
	palettes := make([]*Palette, 0, len(nordColors))
	for variant, defs := range nordColors {
		name := "Nord " + variant
		palette := NewPalette(name, "Nord", variant)

		for _, color := range defs {
			palette.AddColor(color.Name, color.Hex)
		}
		palettes = append(palettes, palette)
	}
	return palettes
}
