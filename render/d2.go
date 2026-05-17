package render

import (
	"fmt"
	"io"
	"strings"

	capstack "github.com/plexusone/capability-stack-spec"
)

// D2Style defines the rendering style for D2 diagrams.
type D2Style string

const (
	// D2StyleDefault renders with dependency arrows and detailed layout.
	D2StyleDefault D2Style = "default"

	// D2StyleGrid renders a clean grid layout suitable for executives.
	// No dependency arrows, capabilities aligned in grid within layers.
	D2StyleGrid D2Style = "grid"
)

// D2Options configures D2 diagram generation.
type D2Options struct {
	// Title shown at the top of the diagram.
	Title string

	// Style controls the overall rendering approach.
	Style D2Style

	// ShowDependencies renders dependency arrows between capabilities.
	// Ignored when Style is D2StyleGrid.
	ShowDependencies bool

	// ShowFoundational includes foundational capabilities in a separate section.
	ShowFoundational bool

	// ShowLegend displays the status color legend.
	ShowLegend bool

	// ColorByStatus uses status-based colors instead of category colors.
	ColorByStatus bool

	// Direction sets the layout direction ("down" or "right").
	Direction string

	// GridColumns sets the number of columns per layer row (0 = auto).
	// Only used when Style is D2StyleGrid.
	GridColumns int
}

// DefaultD2Options returns sensible defaults for D2 rendering.
func DefaultD2Options() D2Options {
	return D2Options{
		Style:            D2StyleDefault,
		ShowDependencies: true,
		ShowFoundational: true,
		ShowLegend:       true,
		ColorByStatus:    true,
		Direction:        "down",
		GridColumns:      0,
	}
}

// GridD2Options returns options optimized for executive grid view.
func GridD2Options() D2Options {
	return D2Options{
		Style:            D2StyleGrid,
		ShowDependencies: false,
		ShowFoundational: true,
		ShowLegend:       true,
		ColorByStatus:    true,
		Direction:        "down",
		GridColumns:      0, // Auto-calculate
	}
}

// statusColors maps capability status to D2 fill colors.
var statusColors = map[string]string{
	capstack.StatusOperational: "#10b981", // green
	capstack.StatusImplemented: "#3b82f6", // blue
	capstack.StatusInProgress:  "#f59e0b", // amber
	capstack.StatusPlanned:     "#9ca3af", // gray
	capstack.StatusDeprecated:  "#ef4444", // red
}

// statusTextColors maps capability status to text colors for contrast.
var statusTextColors = map[string]string{
	capstack.StatusOperational: "#ffffff",
	capstack.StatusImplemented: "#ffffff",
	capstack.StatusInProgress:  "#000000",
	capstack.StatusPlanned:     "#000000",
	capstack.StatusDeprecated:  "#ffffff",
}

// RenderD2 generates a D2 diagram from a CapabilityStack.
func RenderD2(w io.Writer, doc *capstack.CapabilityStack, opts D2Options) error {
	if opts.Style == D2StyleGrid {
		return renderD2Grid(w, doc, opts)
	}
	return renderD2Default(w, doc, opts)
}

// renderD2Default renders the standard view with optional dependency arrows.
func renderD2Default(w io.Writer, doc *capstack.CapabilityStack, opts D2Options) error {
	var b strings.Builder

	// Direction
	if opts.Direction == "right" {
		b.WriteString("direction: right\n\n")
	}

	// Title
	title := opts.Title
	if title == "" && doc.Metadata.Title != "" {
		title = doc.Metadata.Title
	}
	if title != "" {
		b.WriteString(fmt.Sprintf("title: |md\n  # %s\n| {near: top-center}\n\n", title))
	}

	// Build category color map
	categoryColors := buildCategoryColorMap(doc)

	// Render each layer as a container
	for _, layer := range doc.Layers {
		layerCaps := doc.CapabilitiesByLayer(layer.ID)
		if len(layerCaps) == 0 {
			continue
		}

		renderLayerContainer(&b, layer, layerCaps, categoryColors, opts, false)
	}

	// Foundational capabilities (cross-cutting)
	if opts.ShowFoundational && len(doc.Foundational) > 0 {
		renderFoundationalSection(&b, doc.Foundational, categoryColors, opts, false)
	}

	// Dependencies
	if opts.ShowDependencies {
		renderDependencies(&b, doc)
	}

	// Legend
	if opts.ShowLegend {
		renderLegend(&b)
	}

	_, err := io.WriteString(w, b.String())
	return err
}

// renderD2Grid renders a clean grid layout suitable for executives.
func renderD2Grid(w io.Writer, doc *capstack.CapabilityStack, opts D2Options) error {
	var b strings.Builder

	// Title
	title := opts.Title
	if title == "" && doc.Metadata.Title != "" {
		title = doc.Metadata.Title
	}
	if title != "" {
		b.WriteString(fmt.Sprintf("title: |md\n  # %s\n| {near: top-center}\n\n", title))
	}

	// Build category color map
	categoryColors := buildCategoryColorMap(doc)

	// Count layers with capabilities
	layerCount := 0
	for _, layer := range doc.Layers {
		if len(doc.CapabilitiesByLayer(layer.ID)) > 0 {
			layerCount++
		}
	}
	if opts.ShowFoundational && len(doc.Foundational) > 0 {
		layerCount++
	}

	// Outer container to enforce vertical stacking and equal widths
	b.WriteString("stack: {\n")
	b.WriteString(fmt.Sprintf("  grid-rows: %d\n", layerCount))
	b.WriteString("  grid-gap: 16\n")
	b.WriteString("  style.fill: transparent\n")
	b.WriteString("  style.stroke: transparent\n")
	b.WriteString("\n")

	// Render each layer as a grid container inside the stack
	for _, layer := range doc.Layers {
		layerCaps := doc.CapabilitiesByLayer(layer.ID)
		if len(layerCaps) == 0 {
			continue
		}

		renderLayerContainerGrid(&b, layer, layerCaps, categoryColors, opts)
	}

	// Foundational capabilities
	if opts.ShowFoundational && len(doc.Foundational) > 0 {
		renderFoundationalSectionGrid(&b, doc.Foundational, categoryColors, opts)
	}

	b.WriteString("}\n\n")

	// Legend (compact for grid view)
	if opts.ShowLegend {
		renderLegendCompact(&b)
	}

	_, err := io.WriteString(w, b.String())
	return err
}

// renderLayerContainerGrid renders a layer for grid view (inside stack container).
func renderLayerContainerGrid(b *strings.Builder, layer capstack.Layer, caps []capstack.Capability, categoryColors map[string]string, opts D2Options) {
	layerID := sanitizeD2ID(layer.ID)
	b.WriteString(fmt.Sprintf("  %s: %s {\n", layerID, layer.Name))
	b.WriteString("    grid-columns: 6\n")
	b.WriteString("    grid-gap: 8\n")
	b.WriteString("    style.border-radius: 8\n")
	b.WriteString("    style.fill: \"#f1f5f9\"\n")
	b.WriteString("    style.stroke: \"#cbd5e1\"\n")

	// Capabilities within layer
	for _, cap := range caps {
		renderCapabilityGrid(b, cap, categoryColors, opts)
	}

	b.WriteString("  }\n\n")
}

// renderCapabilityGrid renders a single capability for grid view.
func renderCapabilityGrid(b *strings.Builder, cap capstack.Capability, categoryColors map[string]string, opts D2Options) {
	capID := sanitizeD2ID(cap.ID)
	label := cap.Name

	b.WriteString(fmt.Sprintf("    %s: %s {\n", capID, label))
	b.WriteString("      shape: rectangle\n")
	b.WriteString("      style.border-radius: 6\n")
	b.WriteString("      style.stroke-width: 0\n")
	b.WriteString("      style.shadow: true\n")

	// Color
	if opts.ColorByStatus {
		fillColor := statusColors[cap.Status]
		if fillColor == "" {
			fillColor = "#e5e7eb"
		}
		textColor := statusTextColors[cap.Status]
		if textColor == "" {
			textColor = "#000000"
		}
		b.WriteString(fmt.Sprintf("      style.fill: \"%s\"\n", fillColor))
		b.WriteString(fmt.Sprintf("      style.font-color: \"%s\"\n", textColor))
	} else if catColor, ok := categoryColors[cap.CategoryID]; ok {
		b.WriteString(fmt.Sprintf("      style.fill: \"%s\"\n", catColor))
		b.WriteString("      style.font-color: \"#ffffff\"\n")
	}

	b.WriteString("    }\n")
}

// renderFoundationalSectionGrid renders foundational capabilities for grid view.
func renderFoundationalSectionGrid(b *strings.Builder, caps []capstack.Capability, categoryColors map[string]string, opts D2Options) {
	b.WriteString("  foundational: Foundational {\n")
	b.WriteString("    grid-columns: 4\n")
	b.WriteString("    grid-gap: 8\n")
	b.WriteString("    style.border-radius: 8\n")
	b.WriteString("    style.fill: \"#fef3c7\"\n")
	b.WriteString("    style.stroke: \"#f59e0b\"\n")

	for _, cap := range caps {
		renderCapabilityGrid(b, cap, categoryColors, opts)
	}

	b.WriteString("  }\n")
}

// buildCategoryColorMap creates a map of category ID to color.
func buildCategoryColorMap(doc *capstack.CapabilityStack) map[string]string {
	categoryColors := make(map[string]string)
	for _, cat := range doc.Categories {
		if cat.Color != "" {
			categoryColors[cat.ID] = cat.Color
		}
	}
	return categoryColors
}

// renderLayerContainer renders a layer with its capabilities.
func renderLayerContainer(b *strings.Builder, layer capstack.Layer, caps []capstack.Capability, categoryColors map[string]string, opts D2Options, useGrid bool) {
	layerID := sanitizeD2ID(layer.ID)
	b.WriteString(fmt.Sprintf("%s: %s {\n", layerID, layer.Name))

	if useGrid {
		// Grid layout for clean executive view
		b.WriteString("  grid-columns: 6\n")
		b.WriteString("  style.border-radius: 8\n")
		b.WriteString("  style.fill: \"#f1f5f9\"\n")
		b.WriteString("  style.stroke: \"#cbd5e1\"\n")
	} else {
		b.WriteString("  style.border-radius: 8\n")
		b.WriteString("  style.fill: \"#f8fafc\"\n")
	}

	// Capabilities within layer
	for _, cap := range caps {
		renderCapability(b, cap, categoryColors, opts, useGrid)
	}

	b.WriteString("}\n\n")
}

// renderCapability renders a single capability box.
func renderCapability(b *strings.Builder, cap capstack.Capability, categoryColors map[string]string, opts D2Options, useGrid bool) {
	capID := sanitizeD2ID(cap.ID)
	label := cap.Name

	b.WriteString(fmt.Sprintf("  %s: %s {\n", capID, label))
	b.WriteString("    shape: rectangle\n")

	if useGrid {
		// Cleaner styling for grid view
		b.WriteString("    style.border-radius: 6\n")
		b.WriteString("    style.stroke-width: 0\n")
		b.WriteString("    style.shadow: true\n")
	} else {
		b.WriteString("    style.border-radius: 4\n")
	}

	// Color
	if opts.ColorByStatus {
		fillColor := statusColors[cap.Status]
		if fillColor == "" {
			fillColor = "#e5e7eb" // default gray
		}
		textColor := statusTextColors[cap.Status]
		if textColor == "" {
			textColor = "#000000"
		}
		b.WriteString(fmt.Sprintf("    style.fill: \"%s\"\n", fillColor))
		b.WriteString(fmt.Sprintf("    style.font-color: \"%s\"\n", textColor))
	} else if catColor, ok := categoryColors[cap.CategoryID]; ok {
		b.WriteString(fmt.Sprintf("    style.fill: \"%s\"\n", catColor))
		b.WriteString("    style.font-color: \"#ffffff\"\n")
	}

	b.WriteString("  }\n")
}

// renderFoundationalSection renders the foundational capabilities section.
func renderFoundationalSection(b *strings.Builder, caps []capstack.Capability, categoryColors map[string]string, opts D2Options, useGrid bool) {
	b.WriteString("foundational: Foundational {\n")
	b.WriteString("  style.border-radius: 8\n")

	if useGrid {
		b.WriteString("  grid-columns: 4\n")
		b.WriteString("  style.fill: \"#fef3c7\"\n")
		b.WriteString("  style.stroke: \"#f59e0b\"\n")
	} else {
		b.WriteString("  style.fill: \"#fef3c7\"\n")
		b.WriteString("  style.stroke: \"#f59e0b\"\n")
		b.WriteString("  style.stroke-dash: 3\n")
	}

	for _, cap := range caps {
		renderCapability(b, cap, categoryColors, opts, useGrid)
	}

	b.WriteString("}\n\n")
}

// renderDependencies renders dependency arrows between capabilities.
func renderDependencies(b *strings.Builder, doc *capstack.CapabilityStack) {
	for _, cap := range doc.AllCapabilities() {
		for _, depID := range cap.Dependencies {
			dep := doc.GetCapabilityByID(depID)
			if dep == nil {
				continue
			}

			fromPath := getCapabilityPath(doc, cap)
			toPath := getCapabilityPath(doc, *dep)

			if fromPath != "" && toPath != "" {
				b.WriteString(fmt.Sprintf("%s <- %s: {style.stroke-dash: 3}\n", toPath, fromPath))
			}
		}
	}
}

// renderLegend renders the status color legend.
func renderLegend(b *strings.Builder) {
	b.WriteString("\n# Legend\n")
	b.WriteString("legend: {\n")
	b.WriteString("  near: bottom-right\n")
	b.WriteString("  style.fill: \"#ffffff\"\n")
	b.WriteString("  style.stroke: \"#e5e7eb\"\n")
	b.WriteString("  \n")
	b.WriteString("  operational: Operational {style.fill: \"#10b981\"; style.font-color: \"#ffffff\"}\n")
	b.WriteString("  implemented: Implemented {style.fill: \"#3b82f6\"; style.font-color: \"#ffffff\"}\n")
	b.WriteString("  in-progress: In Progress {style.fill: \"#f59e0b\"; style.font-color: \"#000000\"}\n")
	b.WriteString("  planned: Planned {style.fill: \"#9ca3af\"; style.font-color: \"#000000\"}\n")
	b.WriteString("}\n")
}

// renderLegendCompact renders a compact horizontal legend for grid view.
func renderLegendCompact(b *strings.Builder) {
	b.WriteString("\n# Legend\n")
	b.WriteString("legend: {\n")
	b.WriteString("  near: bottom-center\n")
	b.WriteString("  grid-columns: 4\n")
	b.WriteString("  grid-gap: 8\n")
	b.WriteString("  style.fill: \"#ffffff\"\n")
	b.WriteString("  style.stroke: \"#e5e7eb\"\n")
	b.WriteString("  style.border-radius: 8\n")
	b.WriteString("  \n")
	b.WriteString("  operational: Operational {style.fill: \"#10b981\"; style.font-color: \"#ffffff\"; style.border-radius: 4}\n")
	b.WriteString("  implemented: Implemented {style.fill: \"#3b82f6\"; style.font-color: \"#ffffff\"; style.border-radius: 4}\n")
	b.WriteString("  in_progress: In Progress {style.fill: \"#f59e0b\"; style.font-color: \"#000000\"; style.border-radius: 4}\n")
	b.WriteString("  planned: Planned {style.fill: \"#9ca3af\"; style.font-color: \"#000000\"; style.border-radius: 4}\n")
	b.WriteString("}\n")
}

// sanitizeD2ID converts a kebab-case ID to a D2-safe identifier.
func sanitizeD2ID(id string) string {
	// D2 IDs can contain hyphens, but we need to quote if there are special chars
	// For simplicity, replace hyphens with underscores
	return strings.ReplaceAll(id, "-", "_")
}

// getCapabilityPath returns the full D2 path for a capability (layer.capability).
func getCapabilityPath(doc *capstack.CapabilityStack, cap capstack.Capability) string {
	capID := sanitizeD2ID(cap.ID)

	// Check if in foundational
	for _, f := range doc.Foundational {
		if f.ID == cap.ID {
			return "foundational." + capID
		}
	}

	// Otherwise, find the layer
	layerID := sanitizeD2ID(cap.LayerID)
	return layerID + "." + capID
}

// RenderD2String is a convenience function that returns the D2 diagram as a string.
func RenderD2String(doc *capstack.CapabilityStack, opts D2Options) (string, error) {
	var b strings.Builder
	if err := RenderD2(&b, doc, opts); err != nil {
		return "", err
	}
	return b.String(), nil
}
