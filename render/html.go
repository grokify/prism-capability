package render

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	capstack "github.com/grokify/prism-capability"
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

	// Overlays provides additional display data (badges, tooltips) for capabilities.
	// This allows external modules to inject data like maturity levels.
	Overlays OverlayProvider
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

  {{if .ShowLegend}}
  <div class="cs-filters">
    <div class="cs-filter-group">
      <div class="cs-filter-label">Status</div>
      <div class="cs-filter-options">
        <label class="cs-filter-btn" data-filter="status" data-value="operational">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #10b981;"></span>
          <span>Operational</span>
        </label>
        <label class="cs-filter-btn" data-filter="status" data-value="implemented">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #3b82f6;"></span>
          <span>Implemented</span>
        </label>
        <label class="cs-filter-btn" data-filter="status" data-value="in-progress">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #f59e0b;"></span>
          <span>In Progress</span>
        </label>
        <label class="cs-filter-btn" data-filter="status" data-value="planned">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #9ca3af;"></span>
          <span>Planned</span>
        </label>
        <label class="cs-filter-btn" data-filter="status" data-value="deprecated">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #ef4444;"></span>
          <span>Deprecated</span>
        </label>
      </div>
    </div>
    {{if .HasMaturity}}
    <div class="cs-filter-group">
      <div class="cs-filter-label">Maturity Level</div>
      <div class="cs-filter-options">
        <label class="cs-filter-btn cs-maturity-btn" data-filter="maturity" data-value="1">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #ef4444;"></span>
          <span>M1</span>
        </label>
        <label class="cs-filter-btn cs-maturity-btn" data-filter="maturity" data-value="2">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #f59e0b;"></span>
          <span>M2</span>
        </label>
        <label class="cs-filter-btn cs-maturity-btn" data-filter="maturity" data-value="3">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #eab308;"></span>
          <span>M3</span>
        </label>
        <label class="cs-filter-btn cs-maturity-btn" data-filter="maturity" data-value="4">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #22c55e;"></span>
          <span>M4</span>
        </label>
        <label class="cs-filter-btn cs-maturity-btn" data-filter="maturity" data-value="5">
          <input type="checkbox" checked>
          <span class="cs-filter-color" style="background-color: #3b82f6;"></span>
          <span>M5</span>
        </label>
      </div>
    </div>
    {{end}}
    <div class="cs-filter-actions">
      <button class="cs-btn cs-btn-select-all" onclick="selectAll()">Select All</button>
      <button class="cs-btn cs-btn-clear" onclick="clearAll()">Clear All</button>
    </div>
  </div>
  {{end}}

  <div class="cs-stack">
    {{range .Layers}}
    <div class="cs-layer">
      <div class="cs-layer-header">{{.Name}}</div>
      <div class="cs-capabilities">
        {{range .Capabilities}}
        <div class="cs-capability"
             data-status="{{.Status}}"
             data-maturity="{{.MaturityLevel}}"
             data-bg-color="{{.BgColor}}"
             data-text-color="{{.TextColor}}"
             style="background-color: {{.BgColor}}; color: {{.TextColor}};"
             title="{{.Tooltip}}">
          <span class="cs-cap-name">{{.Name}}</span>
          {{if .BadgeText}}<span class="cs-badge" data-badge-bg="{{.BadgeBgColor}}" data-badge-text="{{.BadgeTextColor}}" style="background-color: {{.BadgeBgColor}}; color: {{.BadgeTextColor}};">{{.BadgeText}}</span>{{end}}
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
        <div class="cs-capability"
             data-status="{{.Status}}"
             data-maturity="{{.MaturityLevel}}"
             data-bg-color="{{.BgColor}}"
             data-text-color="{{.TextColor}}"
             style="background-color: {{.BgColor}}; color: {{.TextColor}};"
             title="{{.Tooltip}}">
          <span class="cs-cap-name">{{.Name}}</span>
          {{if .BadgeText}}<span class="cs-badge" data-badge-bg="{{.BadgeBgColor}}" data-badge-text="{{.BadgeTextColor}}" style="background-color: {{.BadgeBgColor}}; color: {{.BadgeTextColor}};">{{.BadgeText}}</span>{{end}}
        </div>
        {{end}}
      </div>
    </div>
    {{end}}
  </div>
</div>
{{if .Standalone}}
<script>
{{.Script}}
</script>
</body>
</html>{{end}}`

const cssStyles = `
html, body {
  margin: 0;
  padding: 0;
  min-height: 100vh;
}
body.dark {
  background: #0f172a;
}
.cs-container {
  --cs-bg: #ffffff;
  --cs-text: #1f2937;
  --cs-border: #e5e7eb;
  --cs-layer-bg: #f8fafc;
  --cs-inactive-bg: #f3f4f6;
  --cs-inactive-text: #9ca3af;
  --cs-inactive-border: #e5e7eb;
  font-family: system-ui, -apple-system, sans-serif;
  background: var(--cs-bg);
  color: var(--cs-text);
  padding: 24px;
  min-height: 100vh;
  box-sizing: border-box;
}
.cs-container.dark {
  --cs-bg: #0f172a;
  --cs-text: #f1f5f9;
  --cs-border: #334155;
  --cs-layer-bg: #1e293b;
  --cs-inactive-bg: #334155;
  --cs-inactive-text: #94a3b8;
  --cs-inactive-border: #475569;
}
.cs-title {
  font-size: 1.75rem;
  font-weight: 700;
  margin: 0 0 24px 0;
  text-align: center;
  letter-spacing: -0.025em;
}
.cs-filters {
  background: var(--cs-layer-bg);
  border: 1px solid var(--cs-border);
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 24px;
}
.cs-filter-group {
  margin-bottom: 16px;
}
.cs-filter-group:last-of-type {
  margin-bottom: 0;
}
.cs-filter-label {
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  opacity: 0.7;
  margin-bottom: 10px;
}
.cs-filter-options {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.cs-filter-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 14px;
  border-radius: 8px;
  font-size: 0.8125rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.15s ease;
  border: 1px solid var(--cs-border);
  background: transparent;
  user-select: none;
}
.cs-filter-btn:hover {
  background: rgba(255, 255, 255, 0.05);
}
.cs-filter-btn input {
  display: none;
}
.cs-filter-btn.inactive {
  opacity: 0.4;
}
.cs-filter-color {
  width: 14px;
  height: 14px;
  border-radius: 4px;
  flex-shrink: 0;
}
.cs-filter-actions {
  display: flex;
  gap: 8px;
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid var(--cs-border);
}
.cs-btn {
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 500;
  cursor: pointer;
  border: 1px solid var(--cs-border);
  background: transparent;
  color: var(--cs-text);
  transition: all 0.15s ease;
}
.cs-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}
.cs-stack {
  display: flex;
  flex-direction: column;
  gap: 20px;
}
.cs-layer {
  background: var(--cs-layer-bg);
  border: 1px solid var(--cs-border);
  border-radius: 12px;
  padding: 20px;
}
.cs-layer-header {
  font-size: 0.875rem;
  font-weight: 600;
  margin-bottom: 12px;
  opacity: 0.8;
}
.cs-capabilities {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  gap: 10px;
}
.cs-capability {
  padding: 14px 16px;
  border-radius: 8px;
  font-size: 0.875rem;
  font-weight: 500;
  text-align: center;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  transition: all 0.2s ease;
}
.cs-capability.filtered-out {
  background-color: var(--cs-inactive-bg) !important;
  color: var(--cs-inactive-text) !important;
  box-shadow: none;
  border: 1px solid var(--cs-inactive-border);
}
.cs-capability.filtered-out .cs-badge {
  background-color: transparent !important;
  color: var(--cs-inactive-text) !important;
  border: 1px solid var(--cs-inactive-border);
}
.cs-cap-name {
  display: block;
  line-height: 1.3;
}
.cs-badge {
  font-size: 0.625rem;
  font-weight: 700;
  padding: 3px 8px;
  border-radius: 4px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}
.cs-foundational {
  background: linear-gradient(135deg, #fef3c7 0%, #fde68a 100%);
  border: 2px dashed #f59e0b;
}
.dark .cs-foundational {
  background: linear-gradient(135deg, #78350f 0%, #92400e 100%);
  border-color: #b45309;
}
`

const jsScript = `
function applyFilters() {
  const statusFilters = [];
  const maturityFilters = [];

  document.querySelectorAll('.cs-filter-btn[data-filter="status"]').forEach(btn => {
    const checkbox = btn.querySelector('input');
    if (checkbox && checkbox.checked) {
      statusFilters.push(btn.dataset.value);
      btn.classList.remove('inactive');
    } else {
      btn.classList.add('inactive');
    }
  });

  document.querySelectorAll('.cs-filter-btn[data-filter="maturity"]').forEach(btn => {
    const checkbox = btn.querySelector('input');
    if (checkbox && checkbox.checked) {
      maturityFilters.push(parseInt(btn.dataset.value));
      btn.classList.remove('inactive');
    } else {
      btn.classList.add('inactive');
    }
  });

  // Check if maturity filter UI exists (it's conditionally rendered)
  const hasMaturityFilterUI = document.querySelectorAll('.cs-filter-btn[data-filter="maturity"]').length > 0;

  document.querySelectorAll('.cs-capability').forEach(cap => {
    const status = cap.dataset.status;
    const maturity = parseInt(cap.dataset.maturity) || 0;

    // Status: must match one of the checked filters (none checked = nothing shows)
    const statusMatch = statusFilters.includes(status);

    // Maturity: if UI not shown, always pass; if item has no maturity (0), always pass;
    // otherwise must match one of the checked filters
    const maturityMatch = !hasMaturityFilterUI || maturity === 0 || maturityFilters.includes(maturity);

    if (statusMatch && maturityMatch) {
      cap.classList.remove('filtered-out');
      cap.style.backgroundColor = cap.dataset.bgColor;
      cap.style.color = cap.dataset.textColor;
      const badge = cap.querySelector('.cs-badge');
      if (badge) {
        badge.style.backgroundColor = badge.dataset.badgeBg;
        badge.style.color = badge.dataset.badgeText;
      }
    } else {
      cap.classList.add('filtered-out');
    }
  });
}

function selectAll() {
  document.querySelectorAll('.cs-filter-btn input').forEach(cb => {
    cb.checked = true;
  });
  applyFilters();
}

function clearAll() {
  document.querySelectorAll('.cs-filter-btn input').forEach(cb => {
    cb.checked = false;
  });
  applyFilters();
}

document.addEventListener('DOMContentLoaded', function() {
  // Use change event on checkboxes instead of click on labels
  // This avoids double-toggle issues with label/checkbox behavior
  document.querySelectorAll('.cs-filter-btn input').forEach(checkbox => {
    checkbox.addEventListener('change', applyFilters);
  });

  // Apply initial filter state
  applyFilters();
});
`

type htmlData struct {
	Title        string
	Styles       template.CSS
	Script       template.JS
	Standalone   bool
	DarkTheme    bool
	ShowLegend   bool
	HasMaturity  bool
	Layers       []htmlLayerData
	Foundational []htmlCapData
}

type htmlLayerData struct {
	Name         string
	Capabilities []htmlCapData
}

type htmlCapData struct {
	Name           string
	Status         string
	MaturityLevel  int
	BgColor        string
	TextColor      string
	Tooltip        string
	BadgeText      string
	BadgeBgColor   string
	BadgeTextColor string
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

	// Check if any overlays have maturity data
	hasMaturity := false
	if opts.Overlays != nil {
		for _, overlay := range opts.Overlays {
			if len(overlay.BadgeText) > 0 && overlay.BadgeText[0] == 'M' {
				hasMaturity = true
				break
			}
		}
	}

	data := htmlData{
		Title:       title,
		Styles:      template.CSS(cssStyles), //nolint:gosec // cssStyles is a constant, not user input
		Script:      template.JS(jsScript),   //nolint:gosec // jsScript is a constant, not user input
		Standalone:  opts.Standalone,
		DarkTheme:   opts.DarkTheme,
		ShowLegend:  opts.ShowLegend,
		HasMaturity: hasMaturity,
		Layers:      make([]htmlLayerData, 0),
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
			layerData.Capabilities[i] = capToHTMLData(cap, opts.Overlays)
		}

		data.Layers = append(data.Layers, layerData)
	}

	// Process foundational
	if opts.ShowFoundational && len(doc.Foundational) > 0 {
		data.Foundational = make([]htmlCapData, len(doc.Foundational))
		for i, cap := range doc.Foundational {
			data.Foundational[i] = capToHTMLData(cap, opts.Overlays)
		}
	}

	return tmpl.Execute(w, data)
}

func capToHTMLData(cap capstack.Capability, overlays OverlayProvider) htmlCapData {
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

	// Get overlay data
	overlay := overlays.Get(cap.ID)
	if overlay.TooltipExtra != "" {
		tooltipParts = append(tooltipParts, overlay.TooltipExtra)
	}

	// Badge colors
	badgeBgColor := overlay.BadgeColor
	if badgeBgColor == "" && overlay.BadgeText != "" {
		badgeBgColor = getBadgeColor(overlay)
	}
	badgeTextColor := overlay.BadgeTextColor
	if badgeTextColor == "" && overlay.BadgeText != "" {
		badgeTextColor = getBadgeTextColor(overlay)
	}

	// Extract maturity level from badge text (e.g., "M3" -> 3, "M2.5" -> 2)
	maturityLevel := 0
	if len(overlay.BadgeText) > 1 && overlay.BadgeText[0] == 'M' {
		// Parse first digit after M
		for i := 1; i < len(overlay.BadgeText); i++ {
			c := overlay.BadgeText[i]
			if c >= '0' && c <= '9' {
				maturityLevel = int(c - '0')
				break
			}
		}
	}

	// Normalize status for data attribute
	status := cap.Status
	if status == "" {
		status = "unknown"
	}

	return htmlCapData{
		Name:           cap.Name,
		Status:         status,
		MaturityLevel:  maturityLevel,
		BgColor:        bgColor,
		TextColor:      textColor,
		Tooltip:        strings.Join(tooltipParts, " | "),
		BadgeText:      overlay.BadgeText,
		BadgeBgColor:   badgeBgColor,
		BadgeTextColor: badgeTextColor,
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
