package capstack

// PRISMRef references a PRISM maturity model for a capability.
type PRISMRef struct {
	// DomainID is the PRISM domain ID (e.g., "security", "operations").
	DomainID string `json:"domainId,omitempty"`

	// SLIIDs lists PRISM SLI IDs that measure this capability's maturity.
	SLIIDs []string `json:"sliIds,omitempty"`

	// LevelCriteria describes what each maturity level means for this capability.
	LevelCriteria *LevelCriteria `json:"levelCriteria,omitempty"`
}

// LevelCriteria defines maturity level descriptions (M1-M5).
type LevelCriteria struct {
	M1 string `json:"M1,omitempty"`
	M2 string `json:"M2,omitempty"`
	M3 string `json:"M3,omitempty"`
	M4 string `json:"M4,omitempty"`
	M5 string `json:"M5,omitempty"`
}

// PRISMIntegration configures global PRISM integration settings.
type PRISMIntegration struct {
	// ModelRef is the path or URL to the PRISM maturity model document.
	ModelRef string `json:"modelRef,omitempty"`

	// StateRef is the path or URL to the PRISM maturity state document.
	StateRef string `json:"stateRef,omitempty"`

	// PlanRef is the path or URL to the PRISM maturity plan document.
	PlanRef string `json:"planRef,omitempty"`

	// DefaultDomain is the default PRISM domain for capabilities without explicit domainId.
	DefaultDomain string `json:"defaultDomain,omitempty"`
}
