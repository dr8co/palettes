package palette

// RosePineColors contains the color definitions for all the Rosé Pine variants.
// Check https://rosepinetheme.com/palette/ingredients for more information.
var RosePineColors = map[string][]ColorDefinition{
	"Base": {
		{"Base", "#191724"},
		{"Surface", "#1f1d2e"},
		{"Overlay", "#26233a"},
		{"Muted", "#6e6a86"},
		{"Subtle", "#908caa"},
		{"Text", "#e0def4"},
		{"Love", "#eb6f92"},
		{"Gold", "#f6c177"},
		{"Rose", "#ebbcba"},
		{"Pine", "#31748f"},
		{"Foam", "#9ccfd8"},
		{"Iris", "#c4a7e7"},
		{"Highlight Low", "#21202e"},
		{"Highlight Med", "#403d52"},
		{"Highlight High", "#524f67"},
	},
	"Moon": {
		{"Base", "#232136"},
		{"Surface", "#2a273f"},
		{"Overlay", "#393552"},
		{"Muted", "#6e6a86"},
		{"Subtle", "#908caa"},
		{"Text", "#e0def4"},
		{"Love", "#eb6f92"},
		{"Gold", "#f6c177"},
		{"Rose", "#ea9a97"},
		{"Pine", "#3e8fb0"},
		{"Foam", "#9ccfd8"},
		{"Iris", "#c4a7e7"},
		{"Highlight Low", "#2a283e"},
		{"Highlight Med", "#44415a"},
		{"Highlight High", "#56526e"},
	},
	"Dawn": {
		{"Base", "#faf4ed"},
		{"Surface", "#fffaf3"},
		{"Overlay", "#f2e9e1"},
		{"Muted", "#9893a5"},
		{"Subtle", "#797593"},
		{"Text", "#575279"},
		{"Love", "#b4637a"},
		{"Gold", "#ea9d34"},
		{"Rose", "#d7827e"},
		{"Pine", "#286983"},
		{"Foam", "#56949f"},
		{"Iris", "#907aa9"},
		{"Highlight Low", "#f4ede8"},
		{"Highlight Med", "#dfdad9"},
		{"Highlight High", "#cecacd"},
	},
}

// CreateRosePinePalette creates a Rosé Pine palette variant with appropriate families.
func CreateRosePinePalette(variant string) *Palette {
	colors, exists := RosePineColors[variant]
	if !exists {
		return nil
	}

	name := "Rose Pine " + variant

	palette := NewPalette(name, "Rose Pine", "dark", variant)

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}
	return palette
}
