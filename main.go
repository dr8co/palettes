package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dr8co/palettes/palette"
	"github.com/dr8co/palettes/registry"
)

const version = "0.1.0"

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
	fmt.Printf("  %s -s dark                      # Show all dark theme palettes (short form)\n", os.Args[0])
	fmt.Printf("  %s -show mocha                  # Show Catppuccin Mocha variant\n", os.Args[0])
	fmt.Printf("  %s -show \"Catppuccin Mocha\"     # Show exact palette name\n", os.Args[0])
	fmt.Printf("  %s -list                        # List all available palettes\n", os.Args[0])
	fmt.Printf("  %s -l                           # List all palettes (short form)\n", os.Args[0])
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = printUsage

	// Define flags with both long and short forms
	showFlag := flags.String("show", "", "Show specific palette or palette family (e.g., 'catppuccin', 'dark', 'mocha')")
	shortShow := flags.String("s", "", "")
	flags.Lookup("s").Usage = flags.Lookup("show").Usage

	listFlag := flags.Bool("list", false, "List all available palettes")
	shortList := flags.Bool("l", false, "")
	flags.Lookup("l").Usage = flags.Lookup("list").Usage

	helpFlag := flags.Bool("help", false, "Show help information")
	shortHelp := flags.Bool("h", false, "")
	flags.Lookup("h").Usage = flags.Lookup("help").Usage

	versionFlag := flags.Bool("version", false, "Show version information")
	shortVersion := flags.Bool("v", false, "")
	flags.Lookup("v").Usage = flags.Lookup("version").Usage

	if err := flags.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	// Handle version flag
	if *versionFlag || *shortVersion {
		fmt.Printf("palettes version %s\n", version)
		return
	}

	// Handle help flag
	if *helpFlag || *shortHelp {
		printUsage()
		return
	}

	// Initialize the registry with all available schemes
	reg := registry.NewSchemeRegistry()
	palette.RegisterAllSchemes(reg)

	// Handle list flag
	if *listFlag || *shortList {
		printPaletteList(reg)
		return
	}

	// Handle show flag
	showValue := *showFlag
	if *shortShow != "" {
		showValue = *shortShow
	}
	if showValue != "" {
		err := handleShowCommand(reg, showValue)
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
