# Capability Lifecycle

This document explains how capabilities progress from roadmap items to mature, operational systems.

## Lifecycle Stages

```
┌──────────────────────────────────────────────────────────────────────────────┐
│                         CAPABILITY LIFECYCLE                                  │
├──────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐       │
│  │ PLANNED │───►│ IN-PROGRESS │───►│ IMPLEMENTED │───►│ OPERATIONAL │       │
│  └─────────┘    └─────────────┘    └─────────────┘    └─────────────┘       │
│       │                                   │                   │              │
│       │                                   │                   │              │
│       ▼                                   ▼                   ▼              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │                           PRISM INTEGRATION                          │    │
│  ├─────────────────────────────────────────────────────────────────────┤    │
│  │  PLANNED/IN-PROGRESS      │  IMPLEMENTED        │  OPERATIONAL      │    │
│  │  ─────────────────────    │  ─────────────      │  ─────────────    │    │
│  │  • In PRISM Plan as       │  • Has PRISM State  │  • Has PRISM State│    │
│  │    initiative             │  • Baseline M1-M2   │  • Level M2-M5    │    │
│  │  • No PRISM State yet     │  • Basic SLI values │  • Full SLI data  │    │
│  │  • Target maturity in     │  • May have gaps    │  • Meets SLOs     │    │
│  │    plan goals             │                     │  • Continuous     │    │
│  │                           │                     │    improvement    │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘
```

## Stage Details

### Planned

The capability is on the roadmap but work has not started.

**In Capability Stack:**
```json
{
  "id": "dast",
  "name": "DAST",
  "status": "planned",
  "priority": "high",
  "targetDate": "2026-07-31",
  "dependencies": ["sast"]
}
```

**In PRISM Plan:**
```json
{
  "initiatives": [{
    "id": "init-implement-dast",
    "name": "Implement DAST scanning",
    "capabilityId": "dast",
    "phase": "q3-2026",
    "targetLevel": 2
  }]
}
```

**Not in PRISM State** - no measurements yet.

---

### In-Progress

Implementation has started but is not yet complete.

**In Capability Stack:**
```json
{
  "id": "threat-modeling",
  "status": "in-progress",
  "targetDate": "2026-06-30",
  "tooling": [
    { "name": "OWASP Threat Dragon", "status": "piloting" }
  ]
}
```

**In PRISM Plan:**
```json
{
  "initiatives": [{
    "id": "init-threat-modeling",
    "status": "in-progress",
    "percentComplete": 40
  }]
}
```

**Not in PRISM State** - still building the capability.

---

### Implemented

The capability exists and is functioning but may not be mature.

**In Capability Stack:**
```json
{
  "id": "sast",
  "status": "implemented",
  "implementedAt": "2026-03-15",
  "tooling": [
    { "name": "Semgrep", "status": "deployed" }
  ]
}
```

**In PRISM State:**
```json
{
  "sliState": {
    "sli-sast-coverage": {
      "qualitativeState": "measured",
      "windows": {
        "30d": { "value": 65, "timestamp": "2026-05-15T00:00:00Z" }
      }
    }
  },
  "maturityState": {
    "sast": {
      "current": { "level": 2, "achievedAt": "2026-03-15" },
      "target": { "level": 4, "targetDate": "2026-09-30" }
    }
  }
}
```

**In PRISM Plan:**
```json
{
  "initiatives": [{
    "id": "init-sast-maturity",
    "name": "Advance SAST to M4",
    "capabilityId": "sast",
    "fromLevel": 2,
    "toLevel": 4
  }]
}
```

---

### Operational

The capability is mature, actively maintained, and meeting SLOs.

**In Capability Stack:**
```json
{
  "id": "incident-response",
  "status": "operational",
  "implementedAt": "2025-06-01"
}
```

**In PRISM State:**
```json
{
  "sliState": {
    "sli-mttr": {
      "qualitativeState": "alerting",
      "windows": {
        "30d": { "value": 25, "timestamp": "2026-05-15T00:00:00Z" }
      }
    }
  },
  "maturityState": {
    "incident-response": {
      "current": { "level": 4, "achievedAt": "2026-01-15" },
      "target": { "level": 5, "targetDate": "2027-06-30" }
    }
  }
}
```

---

### Deprecated

The capability is being phased out in favor of a replacement.

**In Capability Stack:**
```json
{
  "id": "legacy-scanner",
  "status": "deprecated",
  "replacedBy": "sast",
  "sunsetDate": "2026-12-31"
}
```

---

## Greenfield Workflow

When building a new system, most capabilities start as `planned`:

```
Month 1-2: Define capability stack
├── All capabilities marked "planned"
├── Dependencies identified
├── Priorities set (critical → high → medium → low)
└── Target dates assigned

Month 3-6: Foundation phase
├── Critical capabilities → in-progress → implemented
├── PRISM state created for implemented capabilities
├── Baseline measurements captured (M1-M2)
└── High-priority capabilities started

Month 6-12: Build-out phase
├── High/medium capabilities → implemented
├── Critical capabilities → operational (M3-M4)
├── PRISM plan tracks maturity improvements
└── Gap analysis drives prioritization

Year 2+: Optimization phase
├── All core capabilities operational
├── Focus on M4 → M5 maturity
├── New capabilities added as needed
└── Continuous improvement culture
```

## Maturity Progression per Capability

Once a capability is `implemented`, track maturity in PRISM:

```
┌────────────────────────────────────────────────────────────────────┐
│                    MATURITY PROGRESSION                             │
├────────────────────────────────────────────────────────────────────┤
│                                                                    │
│  M1          M2          M3          M4          M5                │
│  ●───────────●───────────●───────────●───────────●                 │
│  │           │           │           │           │                 │
│  │           │           │           │           │                 │
│  ▼           ▼           ▼           ▼           ▼                 │
│  Initial    Basic       Defined     Managed     Optimizing        │
│  ────────   ────────    ────────    ────────    ────────           │
│  • Exists   • Documented• Consistent• Measured  • Continuous       │
│  • Ad-hoc   • Some      • Standards • SLOs met  • Industry-leading │
│  • Manual     process   • Automated • Data-     • Auto-optimizing  │
│              • Reactive   checks      driven    • Predictive       │
│                                                                    │
└────────────────────────────────────────────────────────────────────┘
```

## Example: SAST Capability Journey

| Month | Status | Maturity | Key Actions |
|-------|--------|----------|-------------|
| Jan | planned | - | Added to roadmap, Q2 target |
| Mar | in-progress | - | Semgrep evaluation started |
| Apr | implemented | M2 | Deployed to critical repos |
| Jun | operational | M2 | All repos covered |
| Sep | operational | M3 | CI integration, blocking on high |
| Dec | operational | M4 | Custom rules, <24h remediation |
