package palette

import "github.com/dr8co/palettes/registry"

// RegisterAllSchemes initializes and registers all available color schemes.
func RegisterAllSchemes(reg *registry.SchemeRegistry) {
	// Register Catppuccin palettes
	for variant := range CatppuccinColors {
		if palette := CreateCatppuccinPalette(variant); palette != nil {
			reg.Register(palette)
		}
	}

	// Register Rosé Pine palettes
	for variant := range RosePineColors {
		if palette := CreateRosePinePalette(variant); palette != nil {
			reg.Register(palette)
		}
	}

	// Eldritch palette
	if palette := CreateEldritchPalette(); palette != nil {
		reg.Register(palette)
	}
}
