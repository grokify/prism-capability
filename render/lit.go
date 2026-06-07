package render

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"strings"

	capstack "github.com/grokify/prism-capability"
)

// LitGridData is the JSON structure consumed by the maturity-grid Lit component.
// This matches the TypeScript MaturityGridData interface in prism/ui/src/types.ts.
type LitGridData struct {
	Title        string                 `json:"title,omitempty"`
	Layers       []LitLayer             `json:"layers"`
	Categories   []LitCategory          `json:"categories"`
	Capabilities []LitCapability        `json:"capabilities"`
	Maturity     map[string]LitMaturity `json:"maturity,omitempty"`
}

// LitLayer represents a layer in the Lit component format.
type LitLayer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Order       int    `json:"order,omitempty"`
}

// LitCategory represents a category in the Lit component format.
type LitCategory struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Order       int    `json:"order,omitempty"`
}

// LitCapability represents a capability in the Lit component format.
type LitCapability struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"fullName,omitempty"`
	Description string    `json:"description,omitempty"`
	LayerID     string    `json:"layerId"`
	CategoryID  string    `json:"categoryId,omitempty"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority,omitempty"`
	Importance  string    `json:"importance,omitempty"`
	Order       int       `json:"order,omitempty"`
	Owner       string    `json:"owner,omitempty"`
	Tooling     []LitTool `json:"tooling,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
}

// LitTool represents a tool in the Lit component format.
type LitTool struct {
	Name   string `json:"name"`
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
	URL    string `json:"url,omitempty"`
}

// LitMaturity represents maturity data for a capability.
type LitMaturity struct {
	CapabilityID string `json:"capabilityId"`
	Level        int    `json:"level"`
	SLICount     int    `json:"sliCount,omitempty"`
}

// LitOptions configures Lit component HTML rendering.
type LitOptions struct {
	// Title overrides the document title.
	Title string

	// Theme is "light" or "dark".
	Theme string

	// View is "by-layer" or "by-category".
	View string

	// ShowLegend displays the status/maturity filters.
	ShowLegend bool

	// ShowViewToggle displays the view mode toggle buttons.
	ShowViewToggle bool

	// ComponentPath is the path to the prism-ui.js component file.
	// Defaults to "prism-ui.js" (same directory as HTML).
	ComponentPath string

	// Overlays provides maturity data to include in the JSON.
	Overlays OverlayProvider
}

// DefaultLitOptions returns sensible defaults for Lit rendering.
func DefaultLitOptions() LitOptions {
	return LitOptions{
		Theme:          "light",
		View:           "by-layer",
		ShowLegend:     true,
		ShowViewToggle: true,
		ComponentPath:  "prism-ui.js",
	}
}

// ToLitGridData converts a CapabilityStack to LitGridData format.
func ToLitGridData(doc *capstack.CapabilityStack, opts LitOptions) *LitGridData {
	title := opts.Title
	if title == "" && doc.Metadata.Title != "" {
		title = doc.Metadata.Title
	}

	data := &LitGridData{
		Title:        title,
		Layers:       make([]LitLayer, len(doc.Layers)),
		Categories:   make([]LitCategory, len(doc.Categories)),
		Capabilities: make([]LitCapability, 0, len(doc.Capabilities)+len(doc.Foundational)),
		Maturity:     make(map[string]LitMaturity),
	}

	// Convert layers
	for i, layer := range doc.Layers {
		data.Layers[i] = LitLayer{
			ID:          layer.ID,
			Name:        layer.Name,
			Description: layer.Description,
			Order:       layer.Order,
		}
	}

	// Convert categories
	for i, cat := range doc.Categories {
		data.Categories[i] = LitCategory{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
		}
	}

	// Convert capabilities
	for _, cap := range doc.Capabilities {
		data.Capabilities = append(data.Capabilities, capToLit(cap, opts.Overlays, data.Maturity))
	}

	// Convert foundational capabilities (they still have a layerId even though they span layers)
	for _, cap := range doc.Foundational {
		data.Capabilities = append(data.Capabilities, capToLit(cap, opts.Overlays, data.Maturity))
	}

	// Remove empty maturity map
	if len(data.Maturity) == 0 {
		data.Maturity = nil
	}

	return data
}

func capToLit(cap capstack.Capability, overlays OverlayProvider, maturityMap map[string]LitMaturity) LitCapability {
	status := cap.Status
	if status == "" {
		status = "planned"
	}

	litCap := LitCapability{
		ID:          cap.ID,
		Name:        cap.Name,
		FullName:    cap.FullName,
		Description: cap.Description,
		LayerID:     cap.LayerID,
		CategoryID:  cap.CategoryID,
		Status:      status,
		Priority:    cap.Priority,
		Importance:  cap.Importance,
		Order:       cap.Order,
		Owner:       cap.Owner,
		Tags:        cap.Tags,
	}

	// Convert tooling
	if len(cap.Tooling) > 0 {
		litCap.Tooling = make([]LitTool, len(cap.Tooling))
		for i, tool := range cap.Tooling {
			litCap.Tooling[i] = LitTool{
				Name:   tool.Name,
				Type:   tool.Type,
				Status: tool.Status,
				URL:    tool.URL,
			}
		}
	}

	// Extract maturity from overlay
	if overlays != nil {
		overlay := overlays.Get(cap.ID)
		if overlay.BadgeText != "" && len(overlay.BadgeText) > 1 && overlay.BadgeText[0] == 'M' {
			level := 0
			for i := 1; i < len(overlay.BadgeText); i++ {
				c := overlay.BadgeText[i]
				if c >= '0' && c <= '9' {
					level = int(c - '0')
					break
				}
			}
			if level > 0 {
				maturityMap[cap.ID] = LitMaturity{
					CapabilityID: cap.ID,
					Level:        level,
				}
			}
		}
	}

	return litCap
}

// RenderJSON writes the capability stack as JSON for the Lit component.
func RenderJSON(w io.Writer, doc *capstack.CapabilityStack, opts LitOptions) error {
	data := ToLitGridData(doc, opts)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// RenderJSONString returns the capability stack as a JSON string.
func RenderJSONString(doc *capstack.CapabilityStack, opts LitOptions) (string, error) {
	var b strings.Builder
	if err := RenderJSON(&b, doc, opts); err != nil {
		return "", err
	}
	return b.String(), nil
}

const litHTMLTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}}</title>
  <style>
    html, body {
      margin: 0;
      padding: 0;
      min-height: 100vh;
    }
    body.dark {
      background: #0f172a;
    }
    .theme-toggle {
      position: fixed;
      top: 16px;
      right: 16px;
      z-index: 100;
      padding: 8px 16px;
      border-radius: 8px;
      border: 1px solid #e5e7eb;
      background: #ffffff;
      cursor: pointer;
      font-family: system-ui, sans-serif;
    }
    body.dark .theme-toggle {
      background: #1e293b;
      color: #f1f5f9;
      border-color: #334155;
    }
  </style>
</head>
<body{{if .DarkTheme}} class="dark"{{end}}>
  {{if .ShowThemeToggle}}<button class="theme-toggle" onclick="toggleTheme()">Toggle Theme</button>{{end}}

  <maturity-grid id="grid" view="{{.View}}" theme="{{.Theme}}"{{if .ShowLegend}} show-legend{{end}}{{if .ShowViewToggle}} show-view-toggle{{end}}>
    <script type="application/json">
{{.Data}}
    </script>
  </maturity-grid>

  <script type="module" src="{{.ComponentPath}}"></script>
  {{if .ShowThemeToggle}}<script>
    function toggleTheme() {
      const body = document.body;
      const grid = document.getElementById('grid');
      if (body.classList.contains('dark')) {
        body.classList.remove('dark');
        grid.setAttribute('theme', 'light');
      } else {
        body.classList.add('dark');
        grid.setAttribute('theme', 'dark');
      }
    }
  </script>{{end}}
</body>
</html>`

type litHTMLData struct {
	Title           string
	Theme           string
	View            string
	DarkTheme       bool
	ShowLegend      bool
	ShowViewToggle  bool
	ShowThemeToggle bool
	ComponentPath   string
	Data            template.JS
}

// RenderLitHTML generates an HTML page that loads the Lit maturity-grid component.
func RenderLitHTML(w io.Writer, doc *capstack.CapabilityStack, opts LitOptions) error {
	tmpl, err := template.New("lit").Parse(litHTMLTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Generate JSON data
	gridData := ToLitGridData(doc, opts)
	jsonBytes, err := json.MarshalIndent(gridData, "      ", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	title := opts.Title
	if title == "" && doc.Metadata.Title != "" {
		title = doc.Metadata.Title
	}
	if title == "" {
		title = "Capability Stack"
	}

	theme := opts.Theme
	if theme == "" {
		theme = "light"
	}

	view := opts.View
	if view == "" {
		view = "by-layer"
	}

	componentPath := opts.ComponentPath
	if componentPath == "" {
		componentPath = "prism-ui.js"
	}

	data := litHTMLData{
		Title:           title,
		Theme:           theme,
		View:            view,
		DarkTheme:       theme == "dark",
		ShowLegend:      opts.ShowLegend,
		ShowViewToggle:  opts.ShowViewToggle,
		ShowThemeToggle: true, // Always show theme toggle for now
		ComponentPath:   componentPath,
		Data:            template.JS(jsonBytes), //nolint:gosec // jsonBytes is marshaled from trusted data
	}

	return tmpl.Execute(w, data)
}

// RenderLitHTMLString returns the Lit HTML page as a string.
func RenderLitHTMLString(doc *capstack.CapabilityStack, opts LitOptions) (string, error) {
	var b strings.Builder
	if err := RenderLitHTML(&b, doc, opts); err != nil {
		return "", err
	}
	return b.String(), nil
}
