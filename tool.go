package capstack

// Tool represents a tool or product that implements a capability.
type Tool struct {
	// Name is the tool or product name.
	Name string `json:"name"`

	// Vendor is the vendor name (for commercial tools).
	Vendor string `json:"vendor,omitempty"`

	// Type classifies the tool (commercial, open-source, internal, managed-service).
	Type string `json:"type,omitempty"`

	// URL is the tool's website or documentation link.
	URL string `json:"url,omitempty"`

	// Status is the deployment status (evaluating, piloting, deployed, deprecated).
	Status string `json:"status,omitempty"`
}
