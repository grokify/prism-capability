# PRISM Capability

[![Go Report Card](https://goreportcard.com/badge/github.com/grokify/prism-capability)](https://goreportcard.com/report/github.com/grokify/prism-capability)
[![GoDoc](https://pkg.go.dev/badge/github.com/grokify/prism-capability)](https://pkg.go.dev/github.com/grokify/prism-capability)

Part of the PRISM ecosystem. Defines capability stacks with layers, capabilities, and maturity integration.

## Overview

PRISM Capability defines **what capabilities exist** in an organization's technology landscape. It integrates with [prism-intelligence](https://github.com/grokify/prism-intelligence) to track **maturity levels** and [prism-roadmap](https://github.com/grokify/prism-roadmap) for **improvement roadmaps**.

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
│                       prism-roadmap                                 │
│              OKRs, roadmaps, improvement initiatives                │
│       "Move SAST from M2→M4 by Q3 via these initiatives"            │
└─────────────────────────────────────────────────────────────────────┘
```

## Key Features

- **Capability Stack Definition**: JSON-based specification for organizational capabilities
- **Layered Architecture**: Organize capabilities by technology layer (infrastructure, platform, application)
- **Priority Framework**: Tier capabilities by business criticality (Existential, Critical, Strategic, Enhancement)
- **Dynamic Priority**: Calculate P0-P3 priorities from importance weights and maturity gaps
- **Validation**: Comprehensive validation with dependency cycle detection
- **Rendering**: Generate D2 diagrams and interactive HTML visualizations
- **Overlays**: External modules can inject display data (badges, tooltips) into renderers

## Quick Start

```bash
# Install CLI
go install github.com/grokify/prism-capability/cmd/capstack@latest

# Validate a capability stack
capstack validate my-stack.json

# Render as HTML
capstack render my-stack.json --format html --output stack.html

# List capabilities filtered by status
capstack list my-stack.json --status operational
```

## PRISM Ecosystem

| Package | Purpose |
|---------|---------|
| [prism-core](https://github.com/grokify/prism-core) | Shared types and primitives |
| **prism-capability** | Capability stack definitions |
| [prism-maturity](https://github.com/grokify/prism-maturity) | Maturity models and SLIs |
| [prism-roadmap](https://github.com/grokify/prism-roadmap) | OKRs and improvement plans |
| [prism](https://github.com/grokify/prism) | Unified CLI and site generator |

## Latest Release

**v0.4.0** (2026-05-24)

- Importance field and dynamic priority calculation (P0-P3)
- Priority framework for organizing capabilities into tiers
- Overlay provider for external data injection
- Interactive filters and maturity badges in HTML renderer

See [Release Notes](releases/v0.4.0.md) for details.
