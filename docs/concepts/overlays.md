# Overlays

Overlays allow external modules to inject supplemental display data into capability renderers without tight coupling.

## Overview

The overlay system enables:

- **Maturity badges**: Show M1-M5 levels on capabilities
- **Custom tooltips**: Add context from external data sources
- **Color coding**: Highlight capabilities based on external state

## CapabilityOverlay

```go
type CapabilityOverlay struct {
    // BadgeText is displayed as a badge (e.g., "M3", "L2")
    BadgeText string

    // BadgeColor is the background color (e.g., "#22c55e")
    BadgeColor string

    // BadgeTextColor is the text color
    BadgeTextColor string

    // TooltipExtra is appended to the capability tooltip
    TooltipExtra string
}
```

## OverlayProvider

Map capability IDs to their overlays:

```go
overlays := render.OverlayProvider{
    "sast": {
        BadgeText:    "M3",
        BadgeColor:   "#22c55e",
        TooltipExtra: "Target: M4 by Q3",
    },
    "dast": {
        BadgeText:    "M1",
        BadgeColor:   "#ef4444",
        TooltipExtra: "Critical gap - immediate action",
    },
    "logging": {
        BadgeText:    "M4",
        BadgeColor:   "#3b82f6",
        TooltipExtra: "Meeting SLOs",
    },
}
```

## Using with Renderers

### D2 Diagrams

```go
opts := render.DefaultD2Options()
opts.Overlays = overlays

d2Output, err := render.RenderD2String(cs, opts)
```

Badges appear as styled spans within capability labels:

```d2
sast: |md
  SAST
  <span style="background:#22c55e;">M3</span>
|
```

### HTML Output

```go
opts := render.DefaultHTMLOptions()
opts.Overlays = overlays
opts.ShowLegend = true  // Enables maturity filters

html, err := render.RenderHTMLString(cs, opts)
```

The HTML renderer adds:

- Maturity level filter buttons (M1-M5)
- Badge display on capabilities
- Data attributes for JavaScript filtering

### Lit Output

```go
litOpts := render.LitOptions{
    Theme:      "light",
    View:       "by-layer",
    ShowLegend: true,
}

// Render interactive Lit HTML
var buf bytes.Buffer
render.RenderLitHTML(&buf, cs, litOpts)
```

The Lit renderer provides:

- Interactive filtering and sorting
- View toggle (by-layer / by-category)
- Dark/light theme support
- Modern web component architecture

## Integration with prism-maturity

The primary use case is integrating maturity state from prism-maturity:

```go
import (
    capstack "github.com/grokify/prism-capability"
    "github.com/grokify/prism-capability/render"
    "github.com/grokify/prism-maturity/dashboard"
)

// Load capability stack and maturity state
cs, _ := capstack.LoadFromFile("stack.json")
state, _ := dashboard.LoadState("state.json")

// Build overlays from maturity state
overlays := make(render.OverlayProvider)
for _, cap := range cs.AllCapabilities() {
    if sliIDs := cap.PRISMRef.SliIDs; len(sliIDs) > 0 {
        level := state.GetAggregateLevel(sliIDs)
        overlays[cap.ID] = render.CapabilityOverlay{
            BadgeText:  fmt.Sprintf("M%d", level),
            BadgeColor: dashboard.MaturityColor(level),
        }
    }
}

// Render with overlays
opts := render.DefaultHTMLOptions()
opts.Overlays = overlays
html, _ := render.RenderHTMLString(cs, opts)
```

## Default Badge Colors

If `BadgeColor` is not specified:

- Badge text starting with "M" gets indigo (`#6366f1`)
- Other badges get slate gray (`#64748b`)
- `BadgeTextColor` defaults to white (`#ffffff`)
