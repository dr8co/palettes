//nolint:goconst
package palette

// From https://github.com/ayu-theme/ayu-colors
var ayuColors = map[string][]ColorDefinition{
	"dark": {
		{"gray", "#5a6673"},
		{"red", "#f07178"},
		{"pink", "#f29668"},
		{"orange", "#ff8f40"},
		{"peach", "#e6c08a"},
		{"yellow", "#ffb454"},
		{"green", "#aad94c"},
		{"teal", "#95e6cb"},
		{"indigo", "#39bae6"},
		{"blue", "#59c2ff"},
		{"purple", "#d2a6ff"},
	},
	"light": {
		{"gray", "#adaeb1"},
		{"red", "#f07171"},
		{"pink", "#f2A191"},
		{"orange", "#fa8532"},
		{"peach", "#e59645"},
		{"yellow", "#eba400"},
		{"green", "#86B300"},
		{"teal", "#4cbf99"},
		{"indigo", "#55b4d4"},
		{"blue", "#22a4e6"},
		{"purple", "#a37acc"},
	},
	"mirage": {
		{"gray", "#6e7c8f"},
		{"red", "#f28779"},
		{"pink", "#f29e74"},
		{"orange", "#ffa659"},
		{"peach", "#d9be98"},
		{"yellow", "#ffcd66"},
		{"green", "#d5ff80"},
		{"teal", "#95e6cb"},
		{"indigo", "#5ccfe6"},
		{"blue", "#73d0ff"},
		{"purple", "#dfbfff"},
	},
}

// CreateAyuPalettes creates a set of Ayu palettes.
func CreateAyuPalettes() []*Palette {
	palettes := make([]*Palette, 0, len(ayuColors))
	for nm, defs := range ayuColors {
		name := "Ayu " + nm
		palette := NewPalette(name, "Ayu", "bright", nm)
		if nm == "mirage" {
			palette.AddFamily("dark")
		}

		for _, color := range defs {
			palette.AddColor(color.Name, color.Hex)

			palettes = append(palettes, palette)
		}
	}
	return palettes
}
