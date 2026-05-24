# capstack show

Show details for a specific capability.

## Synopsis

```bash
capstack show <file> <capability-id>
```

## Examples

```bash
capstack show my-stack.json sast
```

## Output

```
Capability: sast
  Name: SAST
  Layer: application
  Category: security
  Status: operational
  Importance: critical
  Order: 1
  Owner: AppSec Team

PRISM Integration:
  Domain: security
  SLI IDs: security-sast-coverage, security-sast-false-positive-rate

Dependencies:
  - ci-cd (required)

Dependents:
  - secure-sdlc
```
