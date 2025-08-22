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
	reg.Register(CreateEldritchPalette())

	// Tokyo Night
	for variant := range TokyoNightColors {
		if palette := CreateTokyoNightPalette(variant); palette != nil {
			reg.Register(palette)
		}
	}

	// Nord
	for variant := range NordColors {
		if palette := CreateNordPalette(variant); palette != nil {
			reg.Register(palette)
		}
	}

	// Gruvbox
	for variant := range GruvboxColors {
		if palette := CreateGruvboxPalette(variant); palette != nil {
			reg.Register(palette)
		}
	}

	// Monokai Pro
	reg.Register(CreateMonokaiProPalette())

	// Dracula
	reg.Register(CreateDraculaPalette())

	// Solarized
	reg.Register(CreateSolarizedPalette())
}
