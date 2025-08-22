package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dr8co/palettes/palette"
	"github.com/dr8co/palettes/registry"
)

// printUsage displays usage information
func printUsage() {
	fmt.Println("Color Palette Display Tool")
	fmt.Println("==========================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("  %s [options]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Printf("  %s                              # Show all palettes\n", os.Args[0])
	fmt.Printf("  %s -show catppuccin             # Show all Catppuccin variants\n", os.Args[0])
	fmt.Printf("  %s -show dark                   # Show all dark theme palettes\n", os.Args[0])
	fmt.Printf("  %s -show mocha                  # Show Catppuccin Mocha variant\n", os.Args[0])
	fmt.Printf("  %s -show \"Catppuccin Mocha\"     # Show exact palette name\n", os.Args[0])
	fmt.Printf("  %s -list                        # List all available palettes\n", os.Args[0])
}

func main() {
	showFlag := flag.String("show", "", "Show specific palette or palette family (e.g., 'catppuccin', 'dark', 'mocha')")
	listFlag := flag.Bool("list", false, "List all available palettes")
	helpFlag := flag.Bool("help", false, "Show help information")

	flag.Usage = printUsage
	flag.Parse()

	// Handle help flag
	if *helpFlag {
		printUsage()
		return
	}

	// Initialize the registry with all available schemes
	reg := registry.NewSchemeRegistry()
	palette.RegisterAllSchemes(reg)

	// Handle list flag
	if *listFlag {
		printPaletteList(reg)
		return
	}

	// Handle show flag
	if *showFlag != "" {
		err := handleShowCommand(reg, *showFlag)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Default: show all palettes
	reg.ShowAll()
}

// printPaletteList displays a list of all available color palettes.
func printPaletteList(reg *registry.SchemeRegistry) {
	fmt.Println("Available color palettes:")
	fmt.Println(strings.Repeat("─", 40))

	for _, name := range reg.List() {
		scheme, _ := reg.Get(name)
		families := strings.Join(scheme.Families(), ", ")
		fmt.Printf("  • %-30s [%s]\n", name, families)
	}

	fmt.Println()
	fmt.Println("Available families:")
	fmt.Println(strings.Repeat("─", 40))

	for _, family := range reg.GetFamilies() {
		count := len(reg.FindByFamily(family))
		fmt.Printf("  • %-15s (%d palette%s)\n", family, count, pluralize(count))
	}
}

// pluralize returns the plural suffix for a given count.
func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

// handleShowCommand processes the '-show' flag to display a specific palette or family.
func handleShowCommand(reg *registry.SchemeRegistry, query string) error {
	query = strings.TrimSpace(strings.ToLower(query))

	// Try an exact match first (case-insensitive)
	for _, name := range reg.List() {
		if strings.ToLower(name) == query {
			if scheme, exists := reg.Get(name); exists {
				scheme.Show()
				return nil
			}
		}
	}

	// Try a family match (e.g., "catppuccin", "dark", "light")
	matches := reg.FindByFamily(query)
	if len(matches) > 0 {
		fmt.Printf("Showing all '%s' palette variants (%d found):\n", query, len(matches))
		fmt.Println(strings.Repeat("═", 60))
		fmt.Println()

		for _, scheme := range matches {
			scheme.Show()
		}
		return nil
	}

	// Try a partial match
	partialMatches := reg.FindByPartialName(query)
	if len(partialMatches) > 0 {
		if len(partialMatches) == 1 {
			partialMatches[0].Show()
			return nil
		}

		fmt.Printf("Multiple palettes match '%s':\n", query)
		for _, scheme := range partialMatches {
			families := strings.Join(scheme.Families(), ", ")
			fmt.Printf("  • %-30s [%s]\n", scheme.Name(), families)
		}
		return fmt.Errorf("please be more specific")
	}

	return fmt.Errorf("no palette found matching '%s'", query)
}
