package capstack

import "strconv"

// PriorityFramework defines the priority tier structure for organizing capabilities.
// This is embedded in the capability stack and can be rendered as an HTML page.
type PriorityFramework struct {
	// Title is the display title for the priority framework page.
	Title string `json:"title,omitempty"`

	// Description explains the purpose of the priority framework.
	Description string `json:"description,omitempty"`

	// Tiers are the priority tiers in order (highest priority first).
	Tiers []PriorityTier `json:"tiers"`
}

// PriorityTier represents a single priority tier grouping capabilities.
type PriorityTier struct {
	// ID is the unique identifier for the tier (e.g., "tier-1").
	ID string `json:"id"`

	// Name is the display name (e.g., "Tier 1").
	Name string `json:"name"`

	// OrderRange defines the capability Order values in this tier.
	OrderRange OrderRange `json:"orderRange"`

	// Orders explains why capabilities in this tier are important.
	// This is the business rationale (e.g., "Existential", "Critical", "Strategic").
	Orders string `json:"orders"`

	// OrdersDescription provides additional context for the Orders rationale.
	OrdersDescription string `json:"ordersDescription,omitempty"`

	// Focus lists the primary capability areas or categories in this tier.
	Focus []string `json:"focus,omitempty"`

	// CapabilityIDs explicitly lists capability IDs in this tier.
	// If provided, overrides OrderRange for determining membership.
	CapabilityIDs []string `json:"capabilityIds,omitempty"`

	// Color is the display color for this tier (hex code).
	Color string `json:"color,omitempty"`
}

// OrderRange defines an inclusive range of Order values.
type OrderRange struct {
	// Min is the minimum Order value (inclusive).
	Min int `json:"min"`

	// Max is the maximum Order value (inclusive). Use 0 for unlimited.
	Max int `json:"max,omitempty"`
}

// Contains returns true if the given order value is within this range.
func (r OrderRange) Contains(order int) bool {
	if order < r.Min {
		return false
	}
	if r.Max > 0 && order > r.Max {
		return false
	}
	return true
}

// String returns a display string for the range (e.g., "1-5" or "31+").
func (r OrderRange) String() string {
	if r.Max == 0 || r.Max < r.Min {
		return strconv.Itoa(r.Min) + "+"
	}
	if r.Min == r.Max {
		return strconv.Itoa(r.Min)
	}
	return strconv.Itoa(r.Min) + "-" + strconv.Itoa(r.Max)
}

// CapabilitiesInTier returns capabilities that belong to this tier.
// If CapabilityIDs is set, uses that list. Otherwise uses OrderRange.
func (t PriorityTier) CapabilitiesInTier(caps []Capability) []Capability {
	var result []Capability

	if len(t.CapabilityIDs) > 0 {
		// Use explicit capability IDs
		idSet := make(map[string]bool)
		for _, id := range t.CapabilityIDs {
			idSet[id] = true
		}
		for _, cap := range caps {
			if idSet[cap.ID] {
				result = append(result, cap)
			}
		}
	} else {
		// Use OrderRange
		for _, cap := range caps {
			if t.OrderRange.Contains(cap.Order) {
				result = append(result, cap)
			}
		}
	}

	return result
}

// GetTierForCapability returns the tier that contains the given capability.
func (pf *PriorityFramework) GetTierForCapability(cap Capability) *PriorityTier {
	for i := range pf.Tiers {
		tier := &pf.Tiers[i]
		// Check explicit IDs first
		if len(tier.CapabilityIDs) > 0 {
			for _, id := range tier.CapabilityIDs {
				if id == cap.ID {
					return tier
				}
			}
		} else if tier.OrderRange.Contains(cap.Order) {
			return tier
		}
	}
	return nil
}

// DefaultPriorityFramework returns a standard 4-tier priority framework.
func DefaultPriorityFramework() *PriorityFramework {
	return &PriorityFramework{
		Title:       "Priority Framework",
		Description: "Capabilities are organized into tiers based on business criticality and strategic importance.",
		Tiers: []PriorityTier{
			{
				ID:                "tier-1",
				Name:              "Tier 1",
				OrderRange:        OrderRange{Min: 1, Max: 5},
				Orders:            "Existential",
				OrdersDescription: "Capabilities essential for business survival and regulatory compliance.",
				Focus:             []string{"Security", "Availability", "Compliance"},
				Color:             "#dc2626", // red-600
			},
			{
				ID:                "tier-2",
				Name:              "Tier 2",
				OrderRange:        OrderRange{Min: 6, Max: 15},
				Orders:            "Critical",
				OrdersDescription: "Capabilities critical for operational excellence and customer trust.",
				Focus:             []string{"Reliability", "Incident Response", "Observability"},
				Color:             "#ea580c", // orange-600
			},
			{
				ID:                "tier-3",
				Name:              "Tier 3",
				OrderRange:        OrderRange{Min: 16, Max: 30},
				Orders:            "Strategic",
				OrdersDescription: "Capabilities that drive competitive advantage and efficiency.",
				Focus:             []string{"Developer Experience", "Automation", "Quality"},
				Color:             "#ca8a04", // yellow-600
			},
			{
				ID:                "tier-4",
				Name:              "Tier 4",
				OrderRange:        OrderRange{Min: 31, Max: 0},
				Orders:            "Enhancement",
				OrdersDescription: "Capabilities that improve experience but are not immediately critical.",
				Focus:             []string{"Nice-to-haves", "Future Capabilities"},
				Color:             "#16a34a", // green-600
			},
		},
	}
}
