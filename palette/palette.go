// Package palette provides functionality for creating, managing, and displaying color palettes
// in terminal environments. It supports various popular color schemes including Catppuccin,
// Dracula, Gruvbox, Monokai, Nord, Rose Pine, Solarized, Tokyo Night, and Eldritch.
//
// The package offers a flexible system for defining and organizing color schemes through
// the [ColorScheme] interface and [Palette] type. Each palette can have multiple colors and
// belong to multiple families (e.g., "dark", "light", "pastel", etc.).
//
// Key features:
//   - Create and manage color palettes with names and families
//   - Add individual colors or multiple colors at once
//   - Display palettes in the terminal with variously styled text
//   - Group palettes by families for better organization
//   - Support for various popular terminal and editor color schemes
//
// Example usage:
//
//	palette := NewPalette("My Theme", "dark", "custom")
//	palette.AddColor("background", "#282a36")
//	palette.AddColor("foreground", "#f8f8f2")
//	palette.Show()
package palette

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	renderer        = lipgloss.NewRenderer(os.Stdout)
	titleCaser      = cases.Title(language.English)
	placeHolderText = [5]string{"Lorem", "ipsum", "dolor", "sit", "amet"}
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
		colors:   make([]Color, 0, 32),
	}
}

// AddColor adds a single color to the palette.
func (p *Palette) AddColor(name, hex string) *Palette {
	color := NewColor(name, hex)
	p.colors = append(p.colors, *color)
	return p
}

// AddColors adds multiple colors to the palette.
func (p *Palette) AddColors(colors ...Color) *Palette {
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

// Show displays the palette.
func (p *Palette) Show() {
	title := renderer.NewStyle().Bold(true).Render(titleCaser.String(p.name))
	fmt.Println("Palette:", title)

	for _, color := range p.colors {
		style := color.Style
		regularText := style.Render(placeHolderText[0])
		italicText := style.Italic(true).Render(placeHolderText[1])
		boldText := style.Bold(true).Render(placeHolderText[2])
		underlineText := style.Underline(true).UnderlineSpaces(true).Render(placeHolderText[3])
		strikethroughText := style.Strikethrough(true).Render(placeHolderText[4])

		bar := style.Reverse(true).Render(fmt.Sprintf(" %-20s %-13s                 ",
			titleCaser.String(color.Def.Name), color.Def.Hex))

		fmt.Printf("%s %s %s %s %s  %s\n", regularText, italicText, boldText, underlineText, strikethroughText, bar)
	}
	fmt.Println(strings.Repeat("─", 80))
	fmt.Println()
}
