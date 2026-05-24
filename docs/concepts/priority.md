# Priority Framework

The priority framework helps organize capabilities by business criticality and calculate dynamic priorities based on maturity gaps.

## Importance Levels

Static importance weights for capabilities, layers, and categories:

| Level | Weight | Description |
|-------|--------|-------------|
| `critical` | 4 | Essential "-ilities" (security, availability, compliance) |
| `high` | 3 | High business impact capabilities |
| `medium` | 2 | Standard importance (default) |
| `low` | 1 | Nice-to-have capabilities |

```json
{
  "id": "sast",
  "name": "SAST",
  "importance": "critical",
  "order": 1
}
```

## Capability Ordering

Explicit ordering for consistent display and tier assignment:

```json
{
  "capabilities": [
    {"id": "sast", "name": "SAST", "order": 1},
    {"id": "dast", "name": "DAST", "order": 2},
    {"id": "sca", "name": "SCA", "order": 3}
  ]
}
```

**Validation rules:**

- If any capability has a non-zero Order, all Order values must be unique
- Capabilities with Order=0 are sorted after those with explicit ordering

## Dynamic Priority (P0-P3)

Calculate dynamic priority from importance and maturity gap:

```
Priority Score = Importance Weight × (Target Level - Current Level)
```

| Score | Priority | Action |
|-------|----------|--------|
| ≥8 | P0 | Immediate action required |
| ≥4 | P1 | High priority improvement |
| ≥2 | P2 | Scheduled improvement |
| <2 | P3 | Low priority enhancement |

### Examples

```go
// Critical capability (weight=4) at M1 with target M3:
// Score = 4 × 2 = 8 → P0
priority := capstack.CalculatePriority("critical", 1, 3)

// Medium capability (weight=2) at M2 with target M3:
// Score = 2 × 1 = 2 → P2
priority := capstack.CalculatePriority("medium", 2, 3)

// Already at target: P3
priority := capstack.CalculatePriority("critical", 3, 3)
```

## Priority Tiers

Organize capabilities into business-priority tiers:

```json
{
  "priorityFramework": {
    "title": "Capability Priorities",
    "description": "Capabilities organized by business criticality",
    "tiers": [
      {
        "id": "tier-1",
        "name": "Tier 1",
        "orderRange": {"min": 1, "max": 5},
        "orders": "Existential",
        "ordersDescription": "Essential for business survival and compliance",
        "focus": ["Security", "Availability", "Compliance"],
        "color": "#dc2626"
      },
      {
        "id": "tier-2",
        "name": "Tier 2",
        "orderRange": {"min": 6, "max": 15},
        "orders": "Critical",
        "ordersDescription": "Critical for operational excellence",
        "focus": ["Reliability", "Incident Response"],
        "color": "#ea580c"
      },
      {
        "id": "tier-3",
        "name": "Tier 3",
        "orderRange": {"min": 16, "max": 30},
        "orders": "Strategic",
        "ordersDescription": "Drive competitive advantage",
        "focus": ["Developer Experience", "Automation"],
        "color": "#ca8a04"
      },
      {
        "id": "tier-4",
        "name": "Tier 4",
        "orderRange": {"min": 31, "max": 0},
        "orders": "Enhancement",
        "ordersDescription": "Nice-to-have improvements",
        "focus": ["Future Capabilities"],
        "color": "#16a34a"
      }
    ]
  }
}
```

### Querying by Tier

```go
pf := cs.PriorityFramework

// Get tier for a capability
tier := pf.GetTierForCapability(cap)

// Get all capabilities in a tier
caps := tier.CapabilitiesInTier(cs.AllCapabilities())
```

## Sorting Methods

Sort capabilities by different criteria:

```go
// By explicit Order field (default)
cs.SortCapabilities(capstack.SortByOrder)

// Alphabetically by name
cs.SortCapabilities(capstack.SortByName)

// By importance weight (critical first)
cs.SortCapabilities(capstack.SortByImportance)

// By static priority field
cs.SortCapabilities(capstack.SortByPriority)

// By lifecycle status (operational first)
cs.SortCapabilities(capstack.SortByStatus)

// Get sorted copy without modifying original
sorted := cs.SortedCapabilities(capstack.SortByImportance)
```
