package main

import (
	"fmt"
	"time"

	capstack "github.com/grokify/prism-capability"
	"github.com/spf13/cobra"
)

var (
	initDomain string
	initOutput string
	initName   string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new capability stack document",
	Long:  `Creates a starter capability stack document with example layers and capabilities.`,
	RunE:  runInit,
}

func init() {
	initCmd.Flags().StringVarP(&initDomain, "domain", "d", "operations", "Primary domain (security, operations, ai, etc.)")
	initCmd.Flags().StringVarP(&initOutput, "output", "o", "capability-stack.json", "Output file path")
	initCmd.Flags().StringVarP(&initName, "name", "n", "", "Stack name (defaults to domain-stack)")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Validate domain
	if err := capstack.ValidateDomain(initDomain); err != nil {
		return fmt.Errorf("invalid domain: %w", err)
	}

	name := initName
	if name == "" {
		name = initDomain + "-stack"
	}

	doc := &capstack.CapabilityStack{
		Schema: "../schema/capability-stack.schema.json",
		Metadata: capstack.Metadata{
			Name:        name,
			Version:     "0.1.0",
			Title:       fmt.Sprintf("%s Capability Stack", toTitle(initDomain)),
			Description: fmt.Sprintf("Capability stack for %s domain", initDomain),
			Domain:      initDomain,
			CreatedAt:   time.Now().Format("2006-01-02"),
			Authors:     []string{"Your Team"},
		},
		Layers: []capstack.Layer{
			{
				ID:          "strategy",
				Name:        "Strategy & Governance",
				Description: "Strategic planning and governance capabilities",
				Order:       1,
				Phase:       "plan",
			},
			{
				ID:          "design",
				Name:        "Design",
				Description: "Design and architecture capabilities",
				Order:       2,
				Phase:       "design",
			},
			{
				ID:          "build",
				Name:        "Build",
				Description: "Build and development capabilities",
				Order:       3,
				Phase:       "build",
			},
			{
				ID:          "operate",
				Name:        "Operate",
				Description: "Operational capabilities",
				Order:       4,
				Phase:       "operate",
			},
		},
		Categories: []capstack.Category{
			{ID: "core", Name: "Core", Color: "#3b82f6"},
			{ID: "automation", Name: "Automation", Color: "#10b981"},
		},
		Capabilities: []capstack.Capability{
			{
				ID:         "example-capability",
				Name:       "Example Capability",
				LayerID:    "strategy",
				CategoryID: "core",
				Status:     capstack.StatusPlanned,
				Priority:   capstack.PriorityMedium,
			},
		},
	}

	if err := doc.SaveToFile(initOutput); err != nil {
		return fmt.Errorf("failed to write %s: %w", initOutput, err)
	}

	fmt.Printf("Created %s\n", initOutput)
	return nil
}

func toTitle(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
