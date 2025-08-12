package palette

import "github.com/dr8co/palettes/registry"

func RegisterAllSchemes(reg *registry.SchemeRegistry) {
	variants := []string{"latte", "frappe", "macchiato", "mocha"}

	for _, variant := range variants {
		if palette := createCatppuccinPalette(variant); palette != nil {
			reg.Register(palette)
		}
	}
}
