package pokeapi

import (
	"fmt"
	"strings"
)

func (p Pokemon) String() string {
	var result strings.Builder

	// Add basic info
	fmt.Fprintf(&result, "Name: %s\nHeight: %d\nWeight: %d\n", p.Name, p.Height, p.Weight)

	// Add stats
	fmt.Fprintln(&result, "Stats:")
	for _, stat := range p.Stats {
		fmt.Fprintf(&result, "  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	// Add types
	fmt.Fprintln(&result, "Types:")
	for _, t := range p.Types {
		fmt.Fprintf(&result, "  - %s\n", t.Type.Name)
	}

	return result.String()
}
