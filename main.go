package main

import (
	"fmt"
	"github.com/dr8co/palettes/palettes"
	"os"
)

// printUsage displays usage information
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run main.go                    # Show all palettes")
	fmt.Println("  go run main.go [palette-name]     # Show specific palette")
	fmt.Println("  go run main.go --list             # List available palettes")
}

func main() {
	registry := palettes.InitAllSchemes()

	if len(os.Args) == 1 {
		// No arguments - show all schemes
		registry.ShowAll()
	} else if len(os.Args) == 2 {
		arg := os.Args[1]

		switch arg {
		case "--help", "-h":
			printUsage()
		case "--list", "-l":
			fmt.Println("Available color palettes:")
			for _, name := range registry.List() {
				fmt.Printf("  • %s\n", name)
			}
		default:
			// Try to show the specific scheme
			err := registry.Show(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Println("\nAvailable palettes:")
				for _, name := range registry.List() {
					fmt.Printf("  • %s\n", name)
				}
				os.Exit(1)
			}
		}
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "Error: Too many arguments\n")
		printUsage()
		os.Exit(1)
	}
}
