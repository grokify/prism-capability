// Package render provides diagram rendering for capability stacks.
package render

// CapabilityOverlay provides additional display data for a capability.
// This allows external modules (like prism-maturity) to inject supplemental
// information without creating tight coupling.
type CapabilityOverlay struct {
	// BadgeText is displayed as a badge on the capability (e.g., "M3.5", "L2").
	BadgeText string

	// BadgeColor is the background color for the badge (e.g., "#3b82f6").
	// If empty, a default color based on the badge content may be used.
	BadgeColor string

	// BadgeTextColor is the text color for the badge.
	// If empty, white or black is chosen based on BadgeColor contrast.
	BadgeTextColor string

	// TooltipExtra is additional text appended to the capability's tooltip.
	TooltipExtra string
}

// OverlayProvider maps capability IDs to their overlay data.
// This is the primary interface for injecting external data into renderers.
type OverlayProvider map[string]CapabilityOverlay

// Get returns the overlay for a capability ID, or an empty overlay if not found.
func (op OverlayProvider) Get(capabilityID string) CapabilityOverlay {
	if op == nil {
		return CapabilityOverlay{}
	}
	return op[capabilityID]
}

// Has returns true if an overlay exists for the given capability ID.
func (op OverlayProvider) Has(capabilityID string) bool {
	if op == nil {
		return false
	}
	_, ok := op[capabilityID]
	return ok
}
