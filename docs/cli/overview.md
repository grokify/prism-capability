# CLI Overview

The `capstack` CLI provides commands for working with capability stack documents.

## Installation

```bash
go install github.com/grokify/prism-capability/cmd/capstack@latest
```

## Commands

| Command | Description |
|---------|-------------|
| `validate` | Validate a capability stack document |
| `init` | Create a new capability stack from template |
| `list` | List capabilities with filtering |
| `show` | Show details for a specific capability |
| `render` | Generate HTML or D2 visualizations |

## Global Flags

| Flag | Description |
|------|-------------|
| `--help`, `-h` | Show help for any command |

## Examples

```bash
# Validate
capstack validate my-stack.json

# List operational capabilities
capstack list my-stack.json --status operational

# Show capability details
capstack show my-stack.json sast

# Render as HTML
capstack render my-stack.json --format html -o stack.html
```
