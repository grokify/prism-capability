# PRISM Capability

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/prism-capability/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/prism-capability/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/prism-capability/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/prism-capability/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/prism-capability/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/prism-capability/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/prism-capability
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/prism-capability
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/prism-capability
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/prism-capability
 [viz-svg]: https://img.shields.io/badge/visualization-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fprism-capability
 [loc-svg]: https://tokei.rs/b1/github/grokify/prism-capability
 [repo-url]: https://github.com/grokify/prism-capability
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/prism-capability/blob/main/LICENSE

Part of the PRISM ecosystem. Defines capability stacks with layers, capabilities, and maturity integration.

## Overview

PRISM Capability defines **what capabilities exist** in an organization's technology landscape. It integrates with [prism-intelligence](https://github.com/grokify/prism-intelligence) to track **maturity levels** and [prism-execution](https://github.com/grokify/prism-execution) for **improvement roadmaps**.

```
┌─────────────────────────────────────────────────────────────────────┐
│                       prism-capability                              │
│              "These are the capabilities we need"                   │
│         Layers, capabilities, categories, relationships             │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              │ references
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│                      prism-intelligence                             │
│         Maturity models, SLIs/SLOs, state tracking                  │
│  "What does M4 look like?" + "Capability X is at M2"                │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              │ references
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│                       prism-execution                               │
│              OKRs, roadmaps, improvement initiatives                │
│       "Move SAST from M2→M4 by Q3 via these initiatives"            │
└─────────────────────────────────────────────────────────────────────┘
```

## Installation

### CLI

```bash
go install github.com/grokify/prism-capability/cmd/capstack@latest
```

### Library

```bash
go get github.com/grokify/prism-capability
```

## CLI Usage

### Validate a document

```bash
capstack validate examples/operations-saas-greenfield.json
```

Output:

```
Valid: examples/operations-saas-greenfield.json
  Name: b2b-saas-operations
  Version: 0.1.0
  Domain: operations
  Layers: 7
  Categories: 6
  Capabilities: 26
  Foundational: 3
```

### List capabilities

```bash
# List all capabilities
capstack list examples/operations-saas-greenfield.json

# Filter by status
capstack list examples/operations-saas-greenfield.json --status=implemented

# Filter by layer
capstack list examples/operations-saas-greenfield.json --layer=observe

# Filter by tag
capstack list examples/security-saas-greenfield.json --tag=shift-left
```

### Show capability details

```bash
capstack show examples/operations-saas-greenfield.json slo-framework
```

Output:

```
ID: slo-framework
Name: SLO Framework
Full Name: Service Level Objectives Framework

Layer: Strategy & Governance (strategy-governance)
Category: Reliability (reliability)

Status: operational
Owner: SRE Team
Implemented At: 2025-03-01

Tooling:
  - Nobl9 (Nobl9) [commercial] - deployed

PRISM Integration:
  Domain: operations
  SLIs: sli-slo-coverage, sli-slo-attainment
  Maturity Criteria:
    M1: No SLOs defined
    M2: SLOs for critical services, manual tracking
    M3: SLOs for all services, automated dashboards
    M4: Error budgets drive release decisions
    M5: Predictive SLO management, auto-scaling based on burn rate
```

### Create a new capability stack

```bash
# Create with defaults
capstack init -o my-stack.json

# Specify domain
capstack init -d security -o security-stack.json

# Specify name and domain
capstack init -n platform-capabilities -d platform -o platform.json
```

### Render to diagram

Generate D2 diagrams or HTML from capability stacks:

```bash
# Generate D2 file
capstack render stack.json -o stack.d2

# Executive grid view (clean, no dependency arrows)
capstack render stack.json --style=grid -o exec-view.d2

# Pipe to D2 CLI for SVG output
capstack render stack.json | d2 - stack.svg
capstack render stack.json --style=grid | d2 - exec.svg

# Generate static HTML (embeddable fragment)
capstack render stack.json -f html

# Generate standalone HTML document
capstack render stack.json -f html --standalone -o stack.html

# Generate dark theme HTML
capstack render stack.json -f html --standalone --dark -o dark.html
```

**Output Formats:**

| Format | Description |
|--------|-------------|
| `d2` | D2 diagram language (default) |
| `html` | Static HTML - embeddable fragment or standalone document |

**Render Styles (D2 only):**

| Style | Description |
|-------|-------------|
| `default` | Standard view with dependency arrows between capabilities |
| `grid` | Clean grid layout for executives - no arrows, aligned boxes, shadows |

**Options:**

| Flag | Description |
|------|-------------|
| `-o, --output` | Output file (default: stdout) |
| `-s, --style` | Render style: `default` or `grid` (D2 only) |
| `-f, --format` | Output format: `d2` or `html` |
| `-t, --title` | Override diagram title |
| `--no-deps` | Hide dependency arrows (D2 only) |
| `--no-foundational` | Hide foundational capabilities |
| `--no-legend` | Hide status legend |
| `--color-by` | Color scheme: `status` (default) or `category` (D2 only) |
| `--standalone` | Generate complete HTML document (HTML only) |
| `--dark` | Use dark theme (HTML only) |

**Status Colors:**

| Status | Color |
|--------|-------|
| operational | Green |
| implemented | Blue |
| in-progress | Amber |
| planned | Gray |
| deprecated | Red |

**Installing D2:**

To render D2 to images, install the D2 CLI:

```bash
# macOS
brew install d2

# Render to SVG
d2 stack.d2 stack.svg

# Render to PNG
d2 stack.d2 stack.png
```

## Web Component

A Lit-based web component is available for embedding capability stacks in websites and MkDocs documentation.

### Installation

```bash
cd web
npm install
npm run build
```

### Usage

```html
<script type="module" src="capability-stack.js"></script>

<!-- Load from JSON file -->
<capability-stack
  src="stack.json"
  show-legend
  show-foundational
  interactive>
</capability-stack>

<!-- Dark theme -->
<capability-stack
  src="stack.json"
  theme="dark"
  show-legend>
</capability-stack>
```

### Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `src` | string | URL to capability stack JSON file |
| `data` | object | Inline JSON data (alternative to src) |
| `theme` | `light` \| `dark` | Color theme (default: light) |
| `show-legend` | boolean | Show status color legend |
| `show-foundational` | boolean | Show foundational capabilities |
| `interactive` | boolean | Enable hover tooltips |
| `filter-status` | string | Filter by capability status |

### MkDocs Integration

The web component works with MkDocs and Material for MkDocs (unlike React components):

```html
<!-- docs/stacks/security.md -->
<script type="module" src="../js/capability-stack.js"></script>

<capability-stack
  src="../data/security-stack.json"
  interactive
  show-legend>
</capability-stack>
```

Copy the built `dist/capability-stack.js` to your MkDocs `docs/js/` directory.

## Go Library Usage

### Loading and validating documents

```go
package main

import (
    "fmt"
    "log"

    capstack "github.com/grokify/prism-capability"
)

func main() {
    // Load from file
    doc, err := capstack.LoadFromFile("capability-stack.json")
    if err != nil {
        log.Fatal(err)
    }

    // Validate
    errs := doc.Validate()
    if errs.HasErrors() {
        for _, e := range errs {
            fmt.Printf("Error: %s\n", e.Error())
        }
        return
    }

    fmt.Printf("Loaded %d capabilities\n", len(doc.AllCapabilities()))
}
```

### Querying capabilities

```go
// Get capability by ID
cap := doc.GetCapabilityByID("sast")
if cap != nil {
    fmt.Printf("Found: %s (%s)\n", cap.Name, cap.Status)
}

// Filter by status
implemented := doc.CapabilitiesByStatus(capstack.StatusImplemented)
fmt.Printf("Implemented: %d\n", len(implemented))

// Filter by layer
buildCaps := doc.CapabilitiesByLayer("build-test")

// Filter by tag
shiftLeft := doc.CapabilitiesByTag("shift-left")

// Get all capabilities including foundational
all := doc.AllCapabilities()
```

### Creating documents programmatically

```go
doc := &capstack.CapabilityStack{
    Metadata: capstack.Metadata{
        Name:    "my-stack",
        Version: "1.0.0",
        Domain:  capstack.DomainOperations,
    },
    Layers: []capstack.Layer{
        {ID: "observe", Name: "Observe", Phase: capstack.PhaseMonitor},
    },
    Capabilities: []capstack.Capability{
        {
            ID:      "metrics",
            Name:    "Metrics Collection",
            LayerID: "observe",
            Status:  capstack.StatusOperational,
            Tooling: []capstack.Tool{
                {Name: "Prometheus", Type: capstack.ToolTypeOpenSource, Status: capstack.ToolStatusDeployed},
            },
        },
    },
}

// Save to file
if err := doc.SaveToFile("my-stack.json"); err != nil {
    log.Fatal(err)
}
```

### Accessing the JSON Schema

```go
import "github.com/grokify/prism-capability/schema"

// Get schema as bytes
schemaBytes := schema.CapabilityStackSchema()

// Get schema as string
schemaStr := schema.CapabilityStackSchemaString()

// Get schema as map
schemaMap, err := schema.CapabilityStackSchemaMap()
```

## Key Concepts

### Layers

Horizontal bands representing phases, domains, or stages. Layers are ordered from top to bottom.

```json
{
  "id": "shift-left-testing",
  "name": "Shift-Left Security Testing",
  "description": "Application security testing in the SDLC",
  "order": 5,
  "phase": "test",
  "nistCsfFunction": "detect"
}
```

### Capabilities

Individual boxes within layers representing specific capabilities. Each capability has a lifecycle status.

```json
{
  "id": "sast",
  "name": "SAST",
  "fullName": "Static Application Security Testing",
  "layerId": "shift-left-testing",
  "categoryId": "appsec",
  "status": "implemented",
  "importance": "critical",
  "order": 1,
  "owner": "AppSec Team",
  "prismRef": {
    "domainId": "security",
    "sliIds": ["sli-sast-coverage", "sli-sast-findings"]
  }
}
```

### Importance

Static importance weights for capabilities, layers, and categories. Used with maturity state to calculate dynamic priority (P0-P3).

| Level | Weight | Description |
|-------|--------|-------------|
| `critical` | 4 | Critical "-ilities" (security, availability) |
| `high` | 3 | High importance capabilities |
| `medium` | 2 | Standard importance (default) |
| `low` | 1 | Nice-to-have capabilities |

### Ordering

Capabilities can have an explicit `order` field for sorting:

```json
{
  "capabilities": [
    {"id": "slo-framework", "name": "SLO Framework", "order": 1, ...},
    {"id": "alerting", "name": "Alerting", "order": 2, ...},
    {"id": "dashboards", "name": "Dashboards", "order": 3, ...}
  ]
}
```

**Validation rules:**

- If any capability has a non-zero Order, all Order values must be unique
- Capabilities with Order=0 are sorted after those with explicit ordering
- When no explicit ordering is set, capabilities sort alphabetically by name

### Sorting

Sort capabilities programmatically:

```go
// Sort methods available
capstack.SortByOrder      // By explicit Order field (default)
capstack.SortByName       // Alphabetically by Name
capstack.SortByImportance // By Importance weight (critical first)
capstack.SortByPriority   // By Priority (critical first)
capstack.SortByStatus     // By Status (operational first)

// Sort in place
cs.SortCapabilities(capstack.SortByImportance)

// Get sorted copy
sorted := cs.SortedCapabilities(capstack.SortByOrder)
```

### Dynamic Priority (P0-P3)

Dynamic priority is calculated from importance and maturity gap:

```
Priority Score = Importance Weight × (Target Level - Current Level)
```

| Score | Priority | Description |
|-------|----------|-------------|
| ≥8 | P0 | Immediate action required |
| ≥4 | P1 | High priority improvement |
| ≥2 | P2 | Scheduled improvement |
| <2 | P3 | Low priority enhancement |

```go
// Calculate dynamic priority
priority := capstack.CalculatePriority("critical", 1, 3) // Returns "P0" (4 × 2 = 8)
priority := capstack.CalculatePriority("medium", 2, 3)   // Returns "P2" (2 × 1 = 2)

// Get priority weight for sorting
weight := capstack.DynamicPriorityWeight("P0") // Returns 4
```
```

### Capability Lifecycle

Capabilities progress through lifecycle stages:

| Status | Description | In prism-intelligence? |
|--------|-------------|------------------------|
| `planned` | On roadmap, not yet started | No - tracked in prism-execution |
| `in-progress` | Currently being implemented | No - tracked in prism-execution |
| `implemented` | Exists but may need maturity improvement | Yes - has M1 baseline |
| `operational` | Running in production, actively maintained | Yes - has M2+ level |
| `deprecated` | Being phased out | Yes - may have declining level |

### Greenfield vs Brownfield

**Greenfield (new system):**

- Most capabilities start as `planned`
- `targetDate` indicates when capability should be implemented
- No prism-intelligence entries until capability reaches `implemented`
- Focus on building foundational capabilities first

**Brownfield (existing system):**

- Most capabilities are `implemented` or `operational`
- prism-intelligence tracks current maturity levels
- prism-execution tracks initiatives to improve maturity
- Focus on maturing existing capabilities

## PRISM Integration

### Reference Structure

Each capability can reference PRISM for maturity tracking:

```json
{
  "id": "sast",
  "prismRef": {
    "domainId": "security",
    "sliIds": ["sli-sast-coverage", "sli-sast-findings"],
    "levelCriteria": {
      "M1": "No SAST",
      "M2": "SAST on critical repos, manual runs",
      "M3": "SAST in CI for all repos, blocking on high severity",
      "M4": "Custom rules, <24h remediation SLA for critical",
      "M5": "ML-enhanced detection, <4h remediation for critical"
    }
  }
}
```

### Document Relationships

| Project | Contains | Example |
|---------|----------|---------|
| **prism-capability** | Capability definitions, lifecycle status | `capability/security-saas.json` |
| **prism-intelligence** | Maturity models, SLIs/SLOs, current state | `intelligence/security-model.json` |
| **prism-execution** | OKRs, roadmaps, improvement initiatives | `execution/security-plan.json` |

### Workflow

1. **Define capabilities** in prism-capability
2. **Define SLIs** in prism-intelligence maturity model (one or more per capability)
3. **Set thresholds** per maturity level (M1-M5) in prism-intelligence
4. **Track state** in prism-intelligence as capabilities become operational
5. **Plan improvements** in prism-execution with initiatives per capability

## Categories

Categories provide visual grouping within layers:

```json
{
  "categories": [
    { "id": "appsec", "name": "AppSec", "color": "#10b981" },
    { "id": "supply-chain", "name": "Supply Chain", "color": "#f59e0b" }
  ]
}
```

## Dependencies

Capabilities can declare dependencies:

```json
{
  "id": "dast",
  "dependencies": ["sast"],
  "enables": ["pentest-automation"]
}
```

## Framework Mappings

Map capabilities to compliance frameworks:

```json
{
  "id": "sbom",
  "frameworkMappings": [
    { "framework": "slsa", "controls": ["L1", "L2"] },
    { "framework": "nist-csf-2.0", "controls": ["ID.SC-4"] }
  ]
}
```

Supported frameworks:

- `nist-csf-2.0` - NIST Cybersecurity Framework 2.0
- `nist-800-53` - NIST SP 800-53
- `iso-27001` - ISO/IEC 27001
- `soc2` - SOC 2 Type II
- `pci-dss` - PCI DSS
- `cis` - CIS Controls
- `mitre-attack` - MITRE ATT&CK
- `owasp` - OWASP Top 10 / ASVS
- `slsa` - SLSA (Supply-chain Levels for Software Artifacts)
- `ssdf` - NIST SSDF (Secure Software Development Framework)

## Validation

The Go library validates:

- **Required fields** - id, name, layerId for capabilities
- **Kebab-case IDs** - All IDs must be kebab-case format
- **Enum values** - Status, priority, importance, domain, phase, etc.
- **Reference integrity** - layerId, categoryId, dependencies reference valid IDs
- **Unique IDs** - No duplicate layer, category, or capability IDs
- **Dependency cycles** - Circular dependencies are detected
- **Order uniqueness** - If any capability has non-zero Order, all Order values must be unique

Example validation errors:

```
capabilities[5].layerId: references non-existent layer ID (value: "invalid-layer")
capabilities[3].status: invalid status "active", must be one of: planned, in-progress, implemented, operational, deprecated
dependencies: circular dependency detected (value: "cap-a -> cap-b -> cap-c -> cap-a")
```

## Examples

- [operations-saas-greenfield.json](examples/operations-saas-greenfield.json) - SRE/Operations reference architecture with SLO framework, observability, incident management
- [security-saas-greenfield.json](examples/security-saas-greenfield.json) - B2B SaaS security stack with AppSec, supply chain, and runtime security

## Schema

See [schema/capability-stack.schema.json](schema/capability-stack.schema.json) for the full JSON Schema.

## Related Projects

| Project | Purpose |
|---------|---------|
| [prism-intelligence](https://github.com/grokify/prism-intelligence) | Maturity models, SLIs/SLOs, state tracking |
| [prism-execution](https://github.com/grokify/prism-execution) | OKRs, roadmaps, improvement initiatives |

## License

MIT
