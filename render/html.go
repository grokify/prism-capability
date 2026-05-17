package render

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	capstack "github.com/plexusone/capability-stack-spec"
)

// HTMLOptions configures static HTML rendering.
type HTMLOptions struct {
	// Title overrides the document title.
	Title string

	// ShowLegend displays the status color legend.
	ShowLegend bool

	// ShowFoundational includes foundational capabilities.
	ShowFoundational bool

	// Standalone generates a complete HTML document (vs. embeddable fragment).
	Standalone bool

	// DarkTheme uses dark color scheme.
	DarkTheme bool
}

// DefaultHTMLOptions returns sensible defaults for HTML rendering.
func DefaultHTMLOptions() HTMLOptions {
	return HTMLOptions{
		ShowLegend:       true,
		ShowFoundational: true,
		Standalone:       false,
		DarkTheme:        false,
	}
}

const htmlTemplate = `{{if .Standalone}}<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}}</title>
  <style>
{{.Styles}}
  </style>
</head>
<body class="{{if .DarkTheme}}dark{{end}}">
{{end}}<div class="cs-container{{if .DarkTheme}} dark{{end}}">
  {{if .Title}}<h1 class="cs-title">{{.Title}}</h1>{{end}}
  <div class="cs-stack">
    {{range .Layers}}
    <div class="cs-layer">
      <div class="cs-layer-header">{{.Name}}</div>
      <div class="cs-capabilities">
        {{range .Capabilities}}
        <div class="cs-capability" style="background-color: {{.BgColor}}; color: {{.TextColor}};" title="{{.Tooltip}}">
          {{.Name}}
        </div>
        {{end}}
      </div>
    </div>
    {{end}}
    {{if .Foundational}}
    <div class="cs-layer cs-foundational">
      <div class="cs-layer-header">Foundational</div>
      <div class="cs-capabilities">
        {{range .Foundational}}
        <div class="cs-capability" style="background-color: {{.BgColor}}; color: {{.TextColor}};" title="{{.Tooltip}}">
          {{.Name}}
        </div>
        {{end}}
      </div>
    </div>
    {{end}}
  </div>
  {{if .ShowLegend}}
  <div class="cs-legend">
    <div class="cs-legend-item"><span class="cs-legend-color" style="background-color: #10b981;"></span>Operational</div>
    <div class="cs-legend-item"><span class="cs-legend-color" style="background-color: #3b82f6;"></span>Implemented</div>
    <div class="cs-legend-item"><span class="cs-legend-color" style="background-color: #f59e0b;"></span>In Progress</div>
    <div class="cs-legend-item"><span class="cs-legend-color" style="background-color: #9ca3af;"></span>Planned</div>
    <div class="cs-legend-item"><span class="cs-legend-color" style="background-color: #ef4444;"></span>Deprecated</div>
  </div>
  {{end}}
</div>{{if .Standalone}}
</body>
</html>{{end}}`

const cssStyles = `
.cs-container {
  --cs-bg: #ffffff;
  --cs-text: #1f2937;
  --cs-border: #e5e7eb;
  --cs-layer-bg: #f8fafc;
  font-family: system-ui, -apple-system, sans-serif;
  background: var(--cs-bg);
  color: var(--cs-text);
  padding: 16px;
}
.cs-container.dark {
  --cs-bg: #1f2937;
  --cs-text: #f9fafb;
  --cs-border: #374151;
  --cs-layer-bg: #111827;
}
.cs-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0 0 16px 0;
  text-align: center;
}
.cs-stack {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.cs-layer {
  background: var(--cs-layer-bg);
  border: 1px solid var(--cs-border);
  border-radius: 8px;
  padding: 16px;
}
.cs-layer-header {
  font-size: 0.875rem;
  font-weight: 600;
  margin-bottom: 8px;
  opacity: 0.8;
}
.cs-capabilities {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 8px;
}
.cs-capability {
  padding: 12px 16px;
  border-radius: 6px;
  font-size: 0.875rem;
  font-weight: 500;
  text-align: center;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}
.cs-foundational {
  background: #fef3c7;
  border: 2px dashed #f59e0b;
}
.dark .cs-foundational {
  background: #78350f;
}
.cs-legend {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 8px;
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid var(--cs-border);
}
.cs-legend-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.75rem;
}
.cs-legend-color {
  width: 16px;
  height: 16px;
  border-radius: 4px;
  display: inline-block;
}
`

type htmlData struct {
	Title        string
	Styles       template.CSS
	Standalone   bool
	DarkTheme    bool
	ShowLegend   bool
	Layers       []htmlLayerData
	Foundational []htmlCapData
}

type htmlLayerData struct {
	Name         string
	Capabilities []htmlCapData
}

type htmlCapData struct {
	Name      string
	BgColor   string
	TextColor string
	Tooltip   string
}

// RenderHTML generates static HTML from a CapabilityStack.
func RenderHTML(w io.Writer, doc *capstack.CapabilityStack, opts HTMLOptions) error {
	tmpl, err := template.New("capstack").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	title := opts.Title
	if title == "" && doc.Metadata.Title != "" {
		title = doc.Metadata.Title
	}

	data := htmlData{
		Title:      title,
		Styles:     template.CSS(cssStyles), //nolint:gosec // cssStyles is a constant, not user input
		Standalone: opts.Standalone,
		DarkTheme:  opts.DarkTheme,
		ShowLegend: opts.ShowLegend,
		Layers:     make([]htmlLayerData, 0),
	}

	// Process layers
	for _, layer := range doc.Layers {
		caps := doc.CapabilitiesByLayer(layer.ID)
		if len(caps) == 0 {
			continue
		}

		layerData := htmlLayerData{
			Name:         layer.Name,
			Capabilities: make([]htmlCapData, len(caps)),
		}

		for i, cap := range caps {
			layerData.Capabilities[i] = capToHTMLData(cap)
		}

		data.Layers = append(data.Layers, layerData)
	}

	// Process foundational
	if opts.ShowFoundational && len(doc.Foundational) > 0 {
		data.Foundational = make([]htmlCapData, len(doc.Foundational))
		for i, cap := range doc.Foundational {
			data.Foundational[i] = capToHTMLData(cap)
		}
	}

	return tmpl.Execute(w, data)
}

func capToHTMLData(cap capstack.Capability) htmlCapData {
	bgColor := statusColors[cap.Status]
	if bgColor == "" {
		bgColor = "#e5e7eb"
	}
	textColor := statusTextColors[cap.Status]
	if textColor == "" {
		textColor = "#000000"
	}

	var tooltipParts []string
	if cap.FullName != "" {
		tooltipParts = append(tooltipParts, cap.FullName)
	}
	if cap.Description != "" {
		tooltipParts = append(tooltipParts, cap.Description)
	}
	if cap.Owner != "" {
		tooltipParts = append(tooltipParts, "Owner: "+cap.Owner)
	}
	if cap.Status != "" {
		tooltipParts = append(tooltipParts, "Status: "+cap.Status)
	}

	return htmlCapData{
		Name:      cap.Name,
		BgColor:   bgColor,
		TextColor: textColor,
		Tooltip:   strings.Join(tooltipParts, " | "),
	}
}

// RenderHTMLString is a convenience function that returns HTML as a string.
func RenderHTMLString(doc *capstack.CapabilityStack, opts HTMLOptions) (string, error) {
	var b strings.Builder
	if err := RenderHTML(&b, doc, opts); err != nil {
		return "", err
	}
	return b.String(), nil
}
