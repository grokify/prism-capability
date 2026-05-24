# API Reference

Full API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/grokify/prism-capability).

## Core Types

### CapabilityStack

The root document type:

```go
type CapabilityStack struct {
    Metadata          Metadata
    Layers            []Layer
    Categories        []Category
    Capabilities      []Capability
    Foundational      []Capability
    PriorityFramework *PriorityFramework
    MarketIntegration *MarketIntegration
}
```

Key methods:

```go
// Load from file
cs, err := capstack.LoadFromFile("stack.json")

// Validate
errs := cs.Validate()

// Query capabilities
byStatus := cs.CapabilitiesByStatus(capstack.StatusOperational)
byLayer := cs.CapabilitiesByLayer("application")
byCategory := cs.CapabilitiesByCategory("security")
all := cs.AllCapabilities()

// Sort
cs.SortCapabilities(capstack.SortByImportance)
sorted := cs.SortedCapabilities(capstack.SortByOrder)
```

### Capability

```go
type Capability struct {
    ID           string
    Name         string
    Description  string
    LayerID      string
    CategoryID   string
    Status       string
    Priority     string
    Importance   string
    Order        int
    TargetDate   string
    Owner        string
    Tags         []string
    Dependencies []string
    PRISMRef     *PRISMRef
    MarketRef    *MarketRef
}
```

### Priority Functions

```go
// Get importance weight (4=critical, 3=high, 2=medium, 1=low)
weight := capstack.ImportanceWeight("critical") // 4

// Calculate dynamic priority (P0-P3)
priority := capstack.CalculatePriority("critical", 1, 3) // "P0"

// Get priority weight
weight := capstack.DynamicPriorityWeight("P0") // 4
```

## Rendering

### HTML

```go
import "github.com/grokify/prism-capability/render"

opts := render.DefaultHTMLOptions()
opts.Standalone = true
opts.DarkTheme = false
opts.ShowLegend = true
opts.Overlays = overlays // optional

html, err := render.RenderHTMLString(cs, opts)
```

### D2 Diagrams

```go
opts := render.DefaultD2Options()
opts.Style = render.D2StyleDefault // or D2StyleGrid
opts.Overlays = overlays // optional

d2, err := render.RenderD2String(cs, opts)
```

### Overlays

```go
overlays := render.OverlayProvider{
    "sast": {
        BadgeText:    "M3",
        BadgeColor:   "#22c55e",
        TooltipExtra: "On track for M4",
    },
}

// Check if overlay exists
if overlays.Has("sast") {
    overlay := overlays.Get("sast")
}
```

## Validation

```go
errs := cs.Validate()
for _, err := range errs {
    fmt.Printf("%s.%s: %s (got: %s)\n",
        err.Field, err.Value, err.Message)
}
```

Validation checks:

- Required fields
- Kebab-case IDs
- Valid enum values
- Reference integrity
- Dependency cycles
- Order uniqueness

## Constants

### Status

```go
const (
    StatusPlanned    = "planned"
    StatusInProgress = "in-progress"
    StatusImplemented = "implemented"
    StatusOperational = "operational"
    StatusDeprecated  = "deprecated"
)
```

### Importance

```go
const (
    ImportanceCritical = "critical"
    ImportanceHigh     = "high"
    ImportanceMedium   = "medium"
    ImportanceLow      = "low"
)
```

### Sort Methods

```go
const (
    SortByOrder      = "order"
    SortByName       = "name"
    SortByImportance = "importance"
    SortByPriority   = "priority"
    SortByStatus     = "status"
)
```
