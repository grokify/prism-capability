package capstack

// Capability represents a single capability in the stack.
type Capability struct {
	// ID is the unique identifier for the capability (kebab-case).
	ID string `json:"id"`

	// Name is the short display name (for diagram boxes).
	Name string `json:"name"`

	// FullName is the expanded name (e.g., "Static Application Security Testing" for "SAST").
	FullName string `json:"fullName,omitempty"`

	// Description explains what this capability provides.
	Description string `json:"description,omitempty"`

	// LayerID references the layer this capability belongs to.
	LayerID string `json:"layerId"`

	// CategoryID references the category for visual grouping.
	CategoryID string `json:"categoryId,omitempty"`

	// Status is the lifecycle status (planned, in-progress, implemented, operational, deprecated).
	Status string `json:"status,omitempty"`

	// Priority is the implementation priority (critical, high, medium, low).
	Priority string `json:"priority,omitempty"`

	// TargetDate is when planned capabilities should be implemented (YYYY-MM-DD).
	TargetDate string `json:"targetDate,omitempty"`

	// ImplementedAt is when the capability was implemented (YYYY-MM-DD).
	ImplementedAt string `json:"implementedAt,omitempty"`

	// Owner is the team or person responsible for this capability.
	Owner string `json:"owner,omitempty"`

	// Tooling lists tools/products implementing this capability.
	Tooling []Tool `json:"tooling,omitempty"`

	// Dependencies lists capability IDs this capability depends on.
	Dependencies []string `json:"dependencies,omitempty"`

	// Enables lists capability IDs that this capability enables.
	Enables []string `json:"enables,omitempty"`

	// Tags are for filtering and classification (kebab-case).
	Tags []string `json:"tags,omitempty"`

	// FrameworkMappings maps to compliance/security framework controls.
	FrameworkMappings []FrameworkMapping `json:"frameworkMappings,omitempty"`

	// PRISMRef links to PRISM maturity model for this capability.
	PRISMRef *PRISMRef `json:"prismRef,omitempty"`
}
