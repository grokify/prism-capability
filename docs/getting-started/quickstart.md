# Quick Start

## Create a Capability Stack

Create a new capability stack document:

```bash
capstack init my-stack.json
```

Or create one manually:

```json
{
  "$schema": "https://raw.githubusercontent.com/grokify/prism-capability/main/schema/capstack.schema.json",
  "metadata": {
    "id": "my-org-capabilities",
    "name": "my-org-capabilities",
    "version": "1.0.0",
    "title": "My Organization Capabilities"
  },
  "layers": [
    {"id": "infrastructure", "name": "Infrastructure", "order": 1},
    {"id": "platform", "name": "Platform", "order": 2},
    {"id": "application", "name": "Application", "order": 3}
  ],
  "categories": [
    {"id": "security", "name": "Security", "color": "#ef4444"},
    {"id": "observability", "name": "Observability", "color": "#3b82f6"}
  ],
  "capabilities": [
    {
      "id": "sast",
      "name": "SAST",
      "layerId": "application",
      "categoryId": "security",
      "status": "implemented",
      "importance": "critical",
      "order": 1
    },
    {
      "id": "logging",
      "name": "Centralized Logging",
      "layerId": "platform",
      "categoryId": "observability",
      "status": "operational",
      "importance": "high",
      "order": 10
    }
  ]
}
```

## Validate

```bash
capstack validate my-stack.json
```

Output:

```
Valid: my-stack.json
  Name: my-org-capabilities
  Version: 1.0.0
  Layers: 3
  Categories: 2
  Capabilities: 2
```

## List Capabilities

```bash
# List all
capstack list my-stack.json

# Filter by status
capstack list my-stack.json --status operational

# Filter by layer
capstack list my-stack.json --layer application
```

## Render Visualizations

```bash
# HTML (interactive)
capstack render my-stack.json --format html --output stack.html

# D2 diagram
capstack render my-stack.json --format d2 --output stack.d2
d2 stack.d2 stack.svg
```

## Use as Library

```go
package main

import (
    "fmt"
    "log"

    capstack "github.com/grokify/prism-capability"
    "github.com/grokify/prism-capability/render"
)

func main() {
    // Load capability stack
    cs, err := capstack.LoadFromFile("my-stack.json")
    if err != nil {
        log.Fatal(err)
    }

    // Validate
    errs := cs.Validate()
    if len(errs) > 0 {
        log.Fatal(errs)
    }

    // Query capabilities
    operational := cs.CapabilitiesByStatus(capstack.StatusOperational)
    fmt.Printf("Operational: %d capabilities\n", len(operational))

    // Sort by importance
    cs.SortCapabilities(capstack.SortByImportance)

    // Render HTML
    html, _ := render.RenderHTMLString(cs, render.DefaultHTMLOptions())
    fmt.Println(html)
}
```
