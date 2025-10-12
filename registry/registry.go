// Package registry provides a centralized management system for color schemes in terminal
// applications. It implements a registry pattern that allows for dynamic registration,
// lookup, and filtering of color schemes.
//
// The registry package serves as the organizational core for managing multiple color schemes,
// providing features such as:
//   - Registration and retrieval of color schemes
//   - Case-sensitive and partial name matching
//   - Family-based grouping and filtering
//   - Sorted listing of available schemes
//   - Bulk operations on registered schemes
//
// Each color scheme in the registry must implement the [ColorScheme] interface, which
// defines the basic operations that any color scheme should support:
//   - Getting the scheme's name
//   - Displaying the scheme in the terminal
//   - Retrieving the scheme's family associations
//
// Example usage:
//
//	reg := registry.NewSchemeRegistry()
//	reg.Register(myColorScheme)
//	schemes := reg.FindByFamily("dark")
//	reg.ShowAll()
//
// The registry ensures consistent access to color schemes throughout the application
// while providing flexible querying capabilities for scheme discovery and management.
package registry

import (
	"fmt"
	"sort"
	"strings"
)

// ColorScheme is re-exported from the palette package to maintain compatibility.
type ColorScheme interface {
	// Name returns the name of the color scheme.
	Name() string

	// Show displays the color scheme.
	Show()

	// Families returns the color families.
	Families() []string
}

// SchemeRegistry manages available color schemes.
type SchemeRegistry struct {
	schemes map[string]ColorScheme
}

// NewSchemeRegistry creates a new scheme registry.
func NewSchemeRegistry() *SchemeRegistry {
	return &SchemeRegistry{
		schemes: make(map[string]ColorScheme),
	}
}

// Register adds a color scheme to the registry.
func (r *SchemeRegistry) Register(scheme ColorScheme) {
	r.schemes[scheme.Name()] = scheme
}

// Get retrieves a color scheme by name (case-sensitive).
func (r *SchemeRegistry) Get(name string) (ColorScheme, bool) {
	scheme, exists := r.schemes[name]
	return scheme, exists
}

// List returns all registered scheme names, sorted alphabetically.
func (r *SchemeRegistry) List() []string {
	names := make([]string, 0, len(r.schemes))
	for name := range r.schemes {
		names = append(names, name)
	}

	sort.Strings(names)
	return names
}

// Show displays a specific color scheme by name.
func (r *SchemeRegistry) Show(name string) error {
	scheme, exists := r.Get(name)
	if !exists {
		return fmt.Errorf("palette %s not found", name)
	}
	scheme.Show()
	return nil
}

// ShowAll displays all registered color schemes.
func (r *SchemeRegistry) ShowAll() {
	for _, name := range r.List() {
		scheme, _ := r.Get(name)
		scheme.Show()
	}
}

// FindByFamily returns all schemes belonging to a specific family.
func (r *SchemeRegistry) FindByFamily(family string) []ColorScheme {
	family = strings.ToLower(family)
	var matches []ColorScheme

	for _, scheme := range r.schemes {
		for _, schemeFamily := range scheme.Families() {
			if strings.ToLower(schemeFamily) == family {
				matches = append(matches, scheme)
				break // Don't add the same scheme multiple times
			}
		}
	}

	// Sort matches by name for a consistent output
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name() < matches[j].Name()
	})

	return matches
}

// FindByPartialName returns all schemes whose names contain the given substring.
func (r *SchemeRegistry) FindByPartialName(partial string) []ColorScheme {
	partial = strings.ToLower(partial)
	var matches []ColorScheme

	for _, scheme := range r.schemes {
		if strings.Contains(strings.ToLower(scheme.Name()), partial) {
			matches = append(matches, scheme)
		}
	}

	// Sort matches by name for a consistent output
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name() < matches[j].Name()
	})

	return matches
}

// GetFamilies returns all unique family names across all schemes.
func (r *SchemeRegistry) GetFamilies() []string {
	familySet := make(map[string]bool)

	for _, scheme := range r.schemes {
		for _, family := range scheme.Families() {
			if family != "" {
				familySet[family] = true
			}
		}
	}

	families := make([]string, 0, len(familySet))
	for family := range familySet {
		families = append(families, family)
	}

	sort.Strings(families)
	return families
}

// GetSchemesByFamilies returns schemes that belong to ALL specified families (AND operation).
func (r *SchemeRegistry) GetSchemesByFamilies(families []string) []ColorScheme {
	if len(families) == 0 {
		return nil
	}

	var matches []ColorScheme

	for _, scheme := range r.schemes {
		hasAllFamilies := true
		schemeFamilies := make(map[string]bool)

		// Build a set of scheme families for a quick lookup
		for _, family := range scheme.Families() {
			schemeFamilies[strings.ToLower(family)] = true
		}

		// Check if a scheme has all required families
		for _, requiredFamily := range families {
			if !schemeFamilies[strings.ToLower(requiredFamily)] {
				hasAllFamilies = false
				break
			}
		}

		if hasAllFamilies {
			matches = append(matches, scheme)
		}
	}

	// Sort matches by name for a consistent output
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name() < matches[j].Name()
	})

	return matches
}
