package main

import (
	"fmt"
	"os"

	capstack "github.com/grokify/prism-capability"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate <file>",
	Short: "Validate a capability stack document",
	Long:  `Validates a capability stack JSON document against the schema and checks for semantic errors.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runValidate,
}

func runValidate(cmd *cobra.Command, args []string) error {
	path := args[0]

	doc, err := capstack.LoadFromFile(path)
	if err != nil {
		return fmt.Errorf("failed to load %s: %w", path, err)
	}

	errs := doc.Validate()
	if errs.HasErrors() {
		fmt.Fprintf(os.Stderr, "Validation failed with %d error(s):\n", len(errs))
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, "  - %s\n", e.Error())
		}
		os.Exit(1)
	}

	fmt.Printf("Valid: %s\n", path)
	fmt.Printf("  Name: %s\n", doc.Metadata.Name)
	fmt.Printf("  Version: %s\n", doc.Metadata.Version)
	fmt.Printf("  Domain: %s\n", doc.Metadata.Domain)
	fmt.Printf("  Layers: %d\n", len(doc.Layers))
	fmt.Printf("  Categories: %d\n", len(doc.Categories))
	fmt.Printf("  Capabilities: %d\n", len(doc.Capabilities))
	fmt.Printf("  Foundational: %d\n", len(doc.Foundational))

	return nil
}
