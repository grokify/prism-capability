package capstack

// MarketRef references market-strategy-engine capabilities that this
// organizational capability enables. This creates traceability from
// internal capabilities to external market competitiveness.
type MarketRef struct {
	// MarketID is the market this capability contributes to (e.g., "security", "crm").
	MarketID string `json:"marketId,omitempty"`

	// CapabilityIDs lists market capability IDs that this org capability enables.
	// These reference capabilities defined in market-strategy-engine Analysis documents.
	CapabilityIDs []string `json:"capabilityIds,omitempty"`

	// Segments indicates which market segments benefit from this capability.
	// Common values: "smb", "mid-market", "enterprise".
	Segments []string `json:"segments,omitempty"`

	// Impact describes how this organizational capability affects market position.
	Impact string `json:"impact,omitempty"`

	// GapContribution describes which market gaps this capability helps close.
	GapContribution []GapContribution `json:"gapContribution,omitempty"`
}

// GapContribution describes how an org capability contributes to closing a market gap.
type GapContribution struct {
	// CapabilityID is the market capability with the gap.
	CapabilityID string `json:"capabilityId"`

	// SegmentID is the target segment for the gap.
	SegmentID string `json:"segmentId,omitempty"`

	// Contribution describes how this org capability helps close the gap.
	// Examples: "primary", "supporting", "enabling".
	Contribution string `json:"contribution,omitempty"`

	// Description explains the relationship.
	Description string `json:"description,omitempty"`
}

// MarketIntegration configures global market-strategy-engine integration settings.
type MarketIntegration struct {
	// AnalysisRef is the path or URL to the market analysis document.
	AnalysisRef string `json:"analysisRef,omitempty"`

	// DefaultMarket is the default market for capabilities without explicit marketId.
	DefaultMarket string `json:"defaultMarket,omitempty"`

	// FocusSegments lists the primary segments the organization is targeting.
	FocusSegments []string `json:"focusSegments,omitempty"`
}
