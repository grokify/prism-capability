package main

import (
	"fmt"
	"strings"

	capstack "github.com/grokify/prism-capability"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <file> <capability-id>",
	Short: "Show details of a capability",
	Long:  `Shows detailed information about a specific capability.`,
	Args:  cobra.ExactArgs(2),
	RunE:  runShow,
}

func runShow(cmd *cobra.Command, args []string) error {
	path := args[0]
	capID := args[1]

	doc, err := capstack.LoadFromFile(path)
	if err != nil {
		return fmt.Errorf("failed to load %s: %w", path, err)
	}

	cap := doc.GetCapabilityByID(capID)
	if cap == nil {
		return fmt.Errorf("capability %q not found", capID)
	}

	fmt.Printf("ID: %s\n", cap.ID)
	fmt.Printf("Name: %s\n", cap.Name)
	if cap.FullName != "" {
		fmt.Printf("Full Name: %s\n", cap.FullName)
	}
	if cap.Description != "" {
		fmt.Printf("Description: %s\n", cap.Description)
	}

	fmt.Println()

	// Layer info
	layer := doc.GetLayerByID(cap.LayerID)
	if layer != nil {
		fmt.Printf("Layer: %s (%s)\n", layer.Name, cap.LayerID)
	} else {
		fmt.Printf("Layer: %s\n", cap.LayerID)
	}

	// Category info
	if cap.CategoryID != "" {
		cat := doc.GetCategoryByID(cap.CategoryID)
		if cat != nil {
			fmt.Printf("Category: %s (%s)\n", cat.Name, cap.CategoryID)
		} else {
			fmt.Printf("Category: %s\n", cap.CategoryID)
		}
	}

	fmt.Println()

	// Status info
	if cap.Status != "" {
		fmt.Printf("Status: %s\n", cap.Status)
	}
	if cap.Priority != "" {
		fmt.Printf("Priority: %s\n", cap.Priority)
	}
	if cap.Owner != "" {
		fmt.Printf("Owner: %s\n", cap.Owner)
	}
	if cap.TargetDate != "" {
		fmt.Printf("Target Date: %s\n", cap.TargetDate)
	}
	if cap.ImplementedAt != "" {
		fmt.Printf("Implemented At: %s\n", cap.ImplementedAt)
	}

	// Tags
	if len(cap.Tags) > 0 {
		fmt.Printf("Tags: %s\n", strings.Join(cap.Tags, ", "))
	}

	// Tooling
	if len(cap.Tooling) > 0 {
		fmt.Println("\nTooling:")
		for _, tool := range cap.Tooling {
			line := fmt.Sprintf("  - %s", tool.Name)
			if tool.Vendor != "" {
				line += fmt.Sprintf(" (%s)", tool.Vendor)
			}
			if tool.Type != "" {
				line += fmt.Sprintf(" [%s]", tool.Type)
			}
			if tool.Status != "" {
				line += fmt.Sprintf(" - %s", tool.Status)
			}
			fmt.Println(line)
		}
	}

	// Dependencies
	if len(cap.Dependencies) > 0 {
		fmt.Println("\nDependencies:")
		for _, depID := range cap.Dependencies {
			dep := doc.GetCapabilityByID(depID)
			if dep != nil {
				fmt.Printf("  - %s (%s)\n", dep.Name, depID)
			} else {
				fmt.Printf("  - %s\n", depID)
			}
		}
	}

	// Enables
	if len(cap.Enables) > 0 {
		fmt.Println("\nEnables:")
		for _, enID := range cap.Enables {
			en := doc.GetCapabilityByID(enID)
			if en != nil {
				fmt.Printf("  - %s (%s)\n", en.Name, enID)
			} else {
				fmt.Printf("  - %s\n", enID)
			}
		}
	}

	// Framework mappings
	if len(cap.FrameworkMappings) > 0 {
		fmt.Println("\nFramework Mappings:")
		for _, fm := range cap.FrameworkMappings {
			fmt.Printf("  - %s: %s\n", fm.Framework, strings.Join(fm.Controls, ", "))
		}
	}

	// PRISM reference
	if cap.PRISMRef != nil {
		fmt.Println("\nPRISM Integration:")
		if cap.PRISMRef.DomainID != "" {
			fmt.Printf("  Domain: %s\n", cap.PRISMRef.DomainID)
		}
		if len(cap.PRISMRef.SLIIDs) > 0 {
			fmt.Printf("  SLIs: %s\n", strings.Join(cap.PRISMRef.SLIIDs, ", "))
		}
		if cap.PRISMRef.LevelCriteria != nil {
			fmt.Println("  Maturity Criteria:")
			lc := cap.PRISMRef.LevelCriteria
			if lc.M1 != "" {
				fmt.Printf("    M1: %s\n", lc.M1)
			}
			if lc.M2 != "" {
				fmt.Printf("    M2: %s\n", lc.M2)
			}
			if lc.M3 != "" {
				fmt.Printf("    M3: %s\n", lc.M3)
			}
			if lc.M4 != "" {
				fmt.Printf("    M4: %s\n", lc.M4)
			}
			if lc.M5 != "" {
				fmt.Printf("    M5: %s\n", lc.M5)
			}
		}
	}

	return nil
}
