// Package cli provides the exported Cobra command tree for the capability stack CLI.
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	capstack "github.com/grokify/prism-capability"
	"github.com/grokify/prism-capability/render"
	"github.com/spf13/cobra"
)

var (
	renderOutput       string
	renderFormat       string
	renderTitle        string
	renderStyle        string
	renderNoDeps       bool
	renderNoFoundation bool
	renderNoLegend     bool
	renderColorBy      string
)

var (
	renderDarkTheme  bool
	renderStandalone bool
)

var renderCmd = &cobra.Command{
	Use:   "render <file>",
	Short: "Render capability stack to a diagram",
	Long: `Renders a capability stack document to a diagram format.

Supported formats:
  d2    D2 diagram language (https://d2lang.com)
  html  Static HTML (embeddable or standalone)

Styles (D2 only):
  default   Standard view with dependency arrows
  grid      Clean grid layout for executives (no arrows, aligned boxes)

Examples:
  prism cap render stack.json -o stack.d2
  prism cap render stack.json --style=grid -o exec-view.d2
  prism cap render stack.json -f d2 | d2 - stack.svg
  prism cap render stack.json --style=grid | d2 - exec.svg
  prism cap render stack.json -f html -o stack.html
  prism cap render stack.json -f html --standalone --dark -o full.html`,
	Args: cobra.ExactArgs(1),
	RunE: runRender,
}

func init() {
	renderCmd.Flags().StringVarP(&renderOutput, "output", "o", "", "Output file (default: stdout)")
	renderCmd.Flags().StringVarP(&renderFormat, "format", "f", "d2", "Output format (d2, html)")
	renderCmd.Flags().StringVarP(&renderTitle, "title", "t", "", "Diagram title (default: from metadata)")
	renderCmd.Flags().StringVarP(&renderStyle, "style", "s", "default", "Render style: default or grid (D2 only)")
	renderCmd.Flags().BoolVar(&renderNoDeps, "no-deps", false, "Hide dependency arrows (D2 only)")
	renderCmd.Flags().BoolVar(&renderNoFoundation, "no-foundational", false, "Hide foundational capabilities")
	renderCmd.Flags().BoolVar(&renderNoLegend, "no-legend", false, "Hide status legend")
	renderCmd.Flags().StringVar(&renderColorBy, "color-by", "status", "Color scheme: status or category (D2 only)")
	renderCmd.Flags().BoolVar(&renderDarkTheme, "dark", false, "Use dark theme (HTML only)")
	renderCmd.Flags().BoolVar(&renderStandalone, "standalone", false, "Generate complete HTML document (HTML only)")
}

func runRender(cmd *cobra.Command, args []string) error {
	path := args[0]

	doc, err := capstack.LoadFromFile(path)
	if err != nil {
		return fmt.Errorf("failed to load %s: %w", path, err)
	}

	// Validate
	errs := doc.Validate()
	if errs.HasErrors() {
		return fmt.Errorf("validation failed: %s", errs.Error())
	}

	// Configure options based on style
	var opts render.D2Options
	switch strings.ToLower(renderStyle) {
	case "grid", "exec", "executive":
		opts = render.GridD2Options()
	default:
		opts = render.DefaultD2Options()
	}

	// Apply overrides
	opts.Title = renderTitle
	opts.ShowDependencies = !renderNoDeps && opts.Style != render.D2StyleGrid
	opts.ShowFoundational = !renderNoFoundation
	opts.ShowLegend = !renderNoLegend
	opts.ColorByStatus = strings.ToLower(renderColorBy) == "status"

	// Determine output
	var out *os.File
	if renderOutput == "" {
		out = os.Stdout
	} else {
		// Auto-detect format from extension if not specified
		if renderFormat == "d2" { // Only auto-detect if using default
			ext := strings.ToLower(filepath.Ext(renderOutput))
			switch ext {
			case ".d2":
				renderFormat = "d2"
			case ".html", ".htm":
				renderFormat = "html"
			}
		}

		var err error
		out, err = os.Create(renderOutput)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer out.Close()
	}

	// Render
	switch strings.ToLower(renderFormat) {
	case "d2":
		if err := render.RenderD2(out, doc, opts); err != nil {
			return fmt.Errorf("render failed: %w", err)
		}
	case "html", "htm":
		htmlOpts := render.HTMLOptions{
			Title:            renderTitle,
			ShowLegend:       !renderNoLegend,
			ShowFoundational: !renderNoFoundation,
			Standalone:       renderStandalone,
			DarkTheme:        renderDarkTheme,
		}
		if err := render.RenderHTML(out, doc, htmlOpts); err != nil {
			return fmt.Errorf("render failed: %w", err)
		}
	default:
		return fmt.Errorf("unsupported format: %s (supported: d2, html)", renderFormat)
	}

	if renderOutput != "" {
		fmt.Fprintf(os.Stderr, "Rendered to %s\n", renderOutput)
	}

	return nil
}
