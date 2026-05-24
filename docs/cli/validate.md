# capstack validate

Validate a capability stack document.

## Synopsis

```bash
capstack validate <file>
```

## Description

Validates the structure and content of a capability stack JSON file, checking:

- Required fields (id, name, layerId for capabilities)
- Kebab-case ID format
- Valid enum values (status, priority, domain)
- Reference integrity (layerId, categoryId exist)
- Dependency cycle detection
- Order uniqueness (if any non-zero Order)

## Examples

```bash
# Validate a document
capstack validate my-stack.json

# Valid output
Valid: my-stack.json
  Name: my-org-capabilities
  Version: 1.0.0
  Domain: operations
  Layers: 7
  Categories: 6
  Capabilities: 26
  Foundational: 3

# Invalid output
Invalid: my-stack.json
Errors:
  - capabilities[5].layerId: references non-existent layer "missing-layer"
  - capabilities[8].order: duplicate order value, also used by capability[2]
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Valid document |
| 1 | Validation errors found |
