//nolint:goconst
package palette

// Contains the color definitions for all the Tokyo Night variants.
// Check https://github.com/tokyo-night/tokyo-night-vscode-theme for more information.
var tokyoNightColors = map[string][]ColorDefinition{
	"dark": {
		{"", "#f7768e"},
		{"", "#ff9e64"},
		{"", "#e0af68"},
		{"", "#cfc9c2"},
		{"", "#9ece6a"},
		{"", "#73daca"},
		{"", "#b4f9f8"},
		{"", "#2ac3de"},
		{"", "#7dcfff"},
		{"", "#7aa2f7"},
		{"", "#bb9af7"},
		{"", "#c0caf5"},
		{"", "#a9b1d6"},
		{"", "#9aa5ce"},
		{"", "#565f89"},
		{"", "#414868"},
		{"", "#24283b"},
		{"", "#1a1b26"},
	},
	"light": {
		{"", "#8c4351"},
		{"", "#965027"},
		{"", "#8f5e15"},
		{"", "#634f30"},
		{"", "#385f0d"},
		{"", "#33635c"},
		{"", "#006c86"},
		{"", "#0f4b6e"},
		{"", "#2959aa"},
		{"", "#5a3e8e"},
		{"", "#343b58"},
		{"", "#40434f"},
		{"", "#343B58"},
		{"", "#6c6e75"},
		{"", "#e6e7ed"},
	},
}

// CreateTokyoNightPalettes creates a set of Tokyo Night palettes.
func CreateTokyoNightPalettes() []*Palette {
	palettes := make([]*Palette, 0, len(tokyoNightColors))
	for variant, defs := range tokyoNightColors {
		name := "Tokyo Night " + variant
		palette := NewPalette(name, "Tokyo Night", "Tokyo", variant)

		for _, color := range defs {
			palette.AddColor(color.Name, color.Hex)
		}
		palettes = append(palettes, palette)
	}
	return palettes
}
