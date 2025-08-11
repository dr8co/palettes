package palettes

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"os"
	"sort"
	"strings"
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
}

type Color struct {
	Name  string
	Hex   string
	Style lipgloss.Style
}

type Palette struct {
	name   string
	colors []Color
}

func NewPalette(name string) *Palette {
	return &Palette{
		name:   name,
		colors: make([]Color, 0, 15),
	}
}

func (p *Palette) AddColor(name, hex string) *Palette {
	p.colors = append(p.colors, Color{Name: name, Hex: hex, Style: createStyle(hex)})
	return p
}

func (p *Palette) AddColors(colors []Color) *Palette {
	p.colors = append(p.colors, colors...)
	return p
}

func (p *Palette) Name() string {
	return p.name
}

func (p *Palette) Colors() []Color {
	return p.colors
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

// SchemeRegistry manages available color schemes.
type SchemeRegistry struct {
	schemes map[string]ColorScheme
}

func NewSchemeRegistry() *SchemeRegistry {
	return &SchemeRegistry{
		schemes: make(map[string]ColorScheme),
	}
}

func (r *SchemeRegistry) Register(scheme ColorScheme) {
	r.schemes[scheme.Name()] = scheme
}

func (r *SchemeRegistry) Get(name string) (ColorScheme, bool) {
	scheme, exists := r.schemes[name]
	return scheme, exists
}

func (r *SchemeRegistry) List() []string {
	names := make([]string, 0, len(r.schemes))
	for name := range r.schemes {
		names = append(names, name)
	}

	sort.Strings(names)
	return names
}

func (r *SchemeRegistry) ShowAll() {
	for _, scheme := range r.schemes {
		scheme.Show()
	}
}

func (r *SchemeRegistry) Show(name string) {
	scheme, exists := r.Get(name)
	if !exists {
		fmt.Println("Scheme not found:", name)
		return
	}
	scheme.Show()
}

type ColorDefinition struct {
	Name string
	Hex  string
}

func createStyle(hex string) lipgloss.Style {
	return renderer.NewStyle().Foreground(lipgloss.Color(hex))
}

func InitAllSchemes() *SchemeRegistry {
	registry := NewSchemeRegistry()
	variants := []string{"latte", "frappe", "macchiato", "mocha"}
	for _, variant := range variants {
		if palette := createCatppuccinPalette(variant); palette != nil {
			registry.Register(palette)
		}
	}

	return registry
}
