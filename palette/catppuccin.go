package palette

// CatppuccinColors contains the color definitions for all the Catppuccin variants
var CatppuccinColors = map[string][]ColorDefinition{
	"latte": {
		{"Rosewater", "#dc8a78"},
		{"Flamingo", "#dd7878"},
		{"Pink", "#ea76cb"},
		{"Mauve", "#8839ef"},
		{"Red", "#d20f39"},
		{"Maroon", "#e64553"},
		{"Peach", "#fe640b"},
		{"Yellow", "#df8e1d"},
		{"Green", "#40a02b"},
		{"Teal", "#179299"},
		{"Sky", "#04a5e5"},
		{"Sapphire", "#209fb5"},
		{"Blue", "#1e66f5"},
		{"Lavender", "#7287fd"},
	},
	"frappe": {
		{"Rosewater", "#f2d5cf"},
		{"Flamingo", "#eebebe"},
		{"Pink", "#f4b8e4"},
		{"Mauve", "#ca9ee6"},
		{"Red", "#e78284"},
		{"Maroon", "#ea999c"},
		{"Peach", "#ef9f76"},
		{"Yellow", "#e5c890"},
		{"Green", "#a6d189"},
		{"Teal", "#81c8be"},
		{"Sky", "#99d1db"},
		{"Sapphire", "#85c1dc"},
		{"Blue", "#8caaee"},
		{"Lavender", "#babbf1"},
	},
	"macchiato": {
		{"Rosewater", "#f4dbd6"},
		{"Flamingo", "#f0c6c6"},
		{"Pink", "#f5bde6"},
		{"Mauve", "#c6a0f6"},
		{"Red", "#ed8796"},
		{"Maroon", "#ee99a0"},
		{"Peach", "#f5a97f"},
		{"Yellow", "#eed49f"},
		{"Green", "#a6da95"},
		{"Teal", "#8bd5ca"},
		{"Sky", "#91d7e3"},
		{"Sapphire", "#7dc4e4"},
		{"Blue", "#8aadf4"},
		{"Lavender", "#b7bdf8"},
	},
	"mocha": {
		{"Rosewater", "#f5e0dc"},
		{"Flamingo", "#f2cdcd"},
		{"Pink", "#f5c2e7"},
		{"Mauve", "#cba6f7"},
		{"Red", "#f38ba8"},
		{"Maroon", "#eba0ac"},
		{"Peach", "#fab387"},
		{"Yellow", "#f9e2af"},
		{"Green", "#a6e3a1"},
		{"Teal", "#94e2d5"},
		{"Sky", "#89dceb"},
		{"Sapphire", "#74c7ec"},
		{"Blue", "#89b4fa"},
		{"Lavender", "#b4befe"},
	},
}

// CatppuccinVariantInfo contains metadata about each Catppuccin variant
var CatppuccinVariantInfo = map[string]struct {
	Theme       string
	Description string
}{
	"latte":     {"light", "Light theme with warm tones"},
	"frappe":    {"dark", "Dark theme with muted colors"},
	"macchiato": {"dark", "Dark theme with vibrant colors"},
	"mocha":     {"dark", "Dark theme with rich colors"},
}

func createCatppuccinPalette(variant string) *Palette {
	colors, exists := CatppuccinColors[variant]
	if !exists {
		return nil
	}

	name := "Catppuccin " + variant

	variantInfo, exists := CatppuccinVariantInfo[variant]
	if !exists {
		return nil
	}

	families := []string{"Catppuccin", variantInfo.Theme, variant}
	palette := NewPalette(name, families...)

	for _, color := range colors {
		palette.AddColor(color.Name, color.Hex)
	}
	return palette
}
