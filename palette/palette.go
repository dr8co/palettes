package palette

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	renderer = lipgloss.NewRenderer(os.Stdout)
)

const (
	placeHolderText = "Lorem ipsum dolor"
)

// ColorScheme defines the interface for color schemes
type ColorScheme interface {
	// Name returns the name of the color scheme.
	Name() string

	// Show displays the color scheme.
	Show()

	// Colors returns the color definitions.
	Colors() []Color

	// Families returns the color families.
	Families() []string
}

// Palette represents a collection of colors with a theme name and families.
type Palette struct {
	name     string
	families []string
	colors   []Color
}

// NewPalette creates a new palette with the given name and families.
func NewPalette(name string, families ...string) *Palette {
	return &Palette{
		name:     name,
		families: families,
		colors:   make([]Color, 0, 15),
	}
}

// AddColor adds a single color to the palette.
func (p *Palette) AddColor(name, hex string) *Palette {
	color := NewColor(name, hex)
	p.colors = append(p.colors, *color)
	return p
}

// AddColors adds multiple colors to the palette.
func (p *Palette) AddColors(colors []Color) *Palette {
	p.colors = append(p.colors, colors...)
	return p
}

// AddFamily adds a family to the palette's families
func (p *Palette) AddFamily(family string) *Palette {
	p.families = append(p.families, family)
	return p
}

// Name returns the palette name.
func (p *Palette) Name() string {
	return p.name
}

// Families returns the palette families.
func (p *Palette) Families() []string {
	return p.families
}

// Colors returns the palette colors.
func (p *Palette) Colors() []Color {
	return p.colors
}

// HasFamily checks if the palette belongs to a specific family.
func (p *Palette) HasFamily(family string) bool {
	family = strings.ToLower(family)
	for _, f := range p.families {
		if strings.ToLower(f) == family {
			return true
		}
	}
	return false
}

// render displays the palette with optional bold styling.
func (p *Palette) render(bold bool) {
	suffix := "Regular"
	if bold {
		suffix = "Bold"
	}
	fmt.Println("Palette:", p.name, suffix)

	for _, color := range p.colors {
		style := color.Style
		if bold {
			style = style.Bold(true)
		}
		bar := style.Reverse(true).Render(fmt.Sprintf(" %-15s %-18s                 ", color.Name, color.Hex))
		fmt.Println(style.Render(fmt.Sprintf("%s %-7s ", placeHolderText, suffix)), bar)
	}
}

// Show displays both regular and bold versions of the palette.
func (p *Palette) Show() {
	// Render regular colors
	p.render(false)
	fmt.Println()

	// Render the bold colors
	p.render(true)
	fmt.Println(strings.Repeat("─", 80))
	fmt.Println()
}
