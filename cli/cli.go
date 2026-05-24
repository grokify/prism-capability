// Package cli provides the exported Cobra command tree for the capability stack CLI.
package cli

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root command for capability stack operations.
// It can be imported and added as a subcommand to other CLI tools.
var RootCmd = &cobra.Command{
	Use:   "capability",
	Short: "PRISM Capability - Capability stacks and dependencies",
	Long: `PRISM Capability provides tools for managing capability stack specifications.

Commands for creating, validating, listing, and rendering capability stacks.`,
}

func init() {
	RootCmd.AddCommand(validateCmd)
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(showCmd)
	RootCmd.AddCommand(renderCmd)
}
