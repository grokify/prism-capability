package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	capstack "github.com/grokify/prism-capability"
	"github.com/spf13/cobra"
)

var (
	listStatus   string
	listLayer    string
	listTag      string
	listCategory string
)

var listCmd = &cobra.Command{
	Use:   "list <file>",
	Short: "List capabilities in a stack",
	Long:  `Lists capabilities with optional filtering by status, layer, category, or tag.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runList,
}

func init() {
	listCmd.Flags().StringVar(&listStatus, "status", "", "Filter by status (planned, in-progress, implemented, operational, deprecated)")
	listCmd.Flags().StringVar(&listLayer, "layer", "", "Filter by layer ID")
	listCmd.Flags().StringVar(&listTag, "tag", "", "Filter by tag")
	listCmd.Flags().StringVar(&listCategory, "category", "", "Filter by category ID")
}

func runList(cmd *cobra.Command, args []string) error {
	path := args[0]

	doc, err := capstack.LoadFromFile(path)
	if err != nil {
		return fmt.Errorf("failed to load %s: %w", path, err)
	}

	caps := doc.AllCapabilities()

	// Apply filters
	if listStatus != "" {
		caps = filterByStatus(caps, listStatus)
	}
	if listLayer != "" {
		caps = filterByLayer(caps, listLayer)
	}
	if listCategory != "" {
		caps = filterByCategory(caps, listCategory)
	}
	if listTag != "" {
		caps = filterByTag(caps, listTag)
	}

	if len(caps) == 0 {
		fmt.Println("No capabilities found matching filters")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "ID\tNAME\tLAYER\tSTATUS\tPRIORITY\tOWNER")
	_, _ = fmt.Fprintln(w, "--\t----\t-----\t------\t--------\t-----")

	for _, cap := range caps {
		status := cap.Status
		if status == "" {
			status = "-"
		}
		priority := cap.Priority
		if priority == "" {
			priority = "-"
		}
		owner := cap.Owner
		if owner == "" {
			owner = "-"
		}

		_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			cap.ID, cap.Name, cap.LayerID, status, priority, owner)
	}

	_ = w.Flush()
	fmt.Printf("\nTotal: %d capabilities\n", len(caps))

	return nil
}

func filterByStatus(caps []capstack.Capability, status string) []capstack.Capability {
	var result []capstack.Capability
	for _, cap := range caps {
		if cap.Status == status {
			result = append(result, cap)
		}
	}
	return result
}

func filterByLayer(caps []capstack.Capability, layerID string) []capstack.Capability {
	var result []capstack.Capability
	for _, cap := range caps {
		if cap.LayerID == layerID {
			result = append(result, cap)
		}
	}
	return result
}

func filterByCategory(caps []capstack.Capability, categoryID string) []capstack.Capability {
	var result []capstack.Capability
	for _, cap := range caps {
		if cap.CategoryID == categoryID {
			result = append(result, cap)
		}
	}
	return result
}

func filterByTag(caps []capstack.Capability, tag string) []capstack.Capability {
	var result []capstack.Capability
	for _, cap := range caps {
		for _, t := range cap.Tags {
			if t == tag {
				result = append(result, cap)
				break
			}
		}
	}
	return result
}
