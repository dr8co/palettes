package palette

import "github.com/dr8co/palettes/registry"

// RegisterAllSchemes initializes and registers all available color schemes.
func RegisterAllSchemes(reg *registry.SchemeRegistry) {
	// Register Catppuccin palettes
	for _, v := range CreateCatppuccinPalettes() {
		reg.Register(v)
	}

	// Register Rosé Pine palettes
	for _, v := range CreateRosePinePalettes() {
		reg.Register(v)
	}

	// Eldritch palette
	reg.Register(CreateEldritchPalette())

	// Tokyo Night
	for _, v := range CreateTokyoNightPalettes() {
		reg.Register(v)
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

	// Everblush
	reg.Register(CreateEverblushPalette())

	// Atom Palettes
	for _, v := range CreateAtomPalettes() {
		reg.Register(v)
	}

	// Ayu Palettes
	for _, v := range CreateAyuPalettes() {
		reg.Register(v)
	}

	// Poimandres
	reg.Register(CreatePoimandresPalette())
}
