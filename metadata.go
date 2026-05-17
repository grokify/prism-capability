package capstack

// Metadata contains document-level information about the capability stack.
type Metadata struct {
	// Name is the identifier for the capability stack (kebab-case).
	Name string `json:"name"`

	// Version is the semantic version of this specification.
	Version string `json:"version"`

	// Title is the display title for rendered output.
	Title string `json:"title,omitempty"`

	// Description provides context about the capability stack.
	Description string `json:"description,omitempty"`

	// Domain is the primary domain (security, ai, platform, etc.).
	Domain string `json:"domain,omitempty"`

	// CreatedAt is the creation date (YYYY-MM-DD format).
	CreatedAt string `json:"createdAt,omitempty"`

	// UpdatedAt is the last update date (YYYY-MM-DD format).
	UpdatedAt string `json:"updatedAt,omitempty"`

	// Authors lists the people/teams who created this stack.
	Authors []string `json:"authors,omitempty"`
}
