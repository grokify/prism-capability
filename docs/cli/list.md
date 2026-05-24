# capstack list

List capabilities with optional filtering.

## Synopsis

```bash
capstack list <file> [flags]
```

## Flags

| Flag | Description |
|------|-------------|
| `--status` | Filter by status (planned, in-progress, implemented, operational, deprecated) |
| `--layer` | Filter by layer ID |
| `--tag` | Filter by tag |

## Examples

```bash
# List all capabilities
capstack list my-stack.json

# Filter by status
capstack list my-stack.json --status operational

# Filter by layer
capstack list my-stack.json --layer application

# Combine filters
capstack list my-stack.json --status implemented --layer platform
```

## Output Format

```
Capabilities (26 total):
  [operational] sast (Application/Security)
  [operational] logging (Platform/Observability)
  [implemented] alerting (Platform/Observability)
  [in-progress] tracing (Platform/Observability)
  [planned] dast (Application/Security)
```
