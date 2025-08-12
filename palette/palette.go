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

type ColorScheme interface {
	Name() string
	Show()
	Colors() []Color
	Families() []string
}

type Palette struct {
	name     string
	families []string
	colors   []Color
}

func NewPalette(name string, families ...string) *Palette {
	return &Palette{
		name:     name,
		families: families,
		colors:   make([]Color, 0, 15),
	}
}

func (p *Palette) AddColor(name, hex string) *Palette {
	color := NewColor(name, hex)
	p.colors = append(p.colors, *color)
	return p
}

func (p *Palette) AddColors(colors []Color) *Palette {
	p.colors = append(p.colors, colors...)
	return p
}

func (p *Palette) AddFamily(family string) *Palette {
	p.families = append(p.families, family)
	return p
}

func (p *Palette) Name() string {
	return p.name
}

func (p *Palette) Families() []string {
	return p.families
}

func (p *Palette) Colors() []Color {
	return p.colors
}

func (p *Palette) HasFamily(family string) bool {
	family = strings.ToLower(family)
	for _, f := range p.families {
		if strings.ToLower(f) == family {
			return true
		}
	}
	return false
}

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

func (p *Palette) Show() {
	// Render regular colors
	p.render(false)
	fmt.Println()

	// Render the bold colors
	p.render(true)
	fmt.Println(strings.Repeat("─", 80))
	fmt.Println()
}
