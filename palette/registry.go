package palette

import "github.com/dr8co/palettes/registry"

// RegisterAllSchemes initializes and registers all available color schemes.
func RegisterAllSchemes(reg *registry.SchemeRegistry) {
	// Register Catppuccin variants
	variants := []string{"latte", "frappe", "macchiato", "mocha"}

	for _, variant := range variants {
		if palette := CreateCatppuccinPalette(variant); palette != nil {
			reg.Register(palette)
		}
	}
}
