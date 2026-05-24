package capstack

// Category groups capabilities within layers for visual organization.
type Category struct {
	// ID is the unique identifier for the category (kebab-case).
	ID string `json:"id"`

	// Name is the display name for the category.
	Name string `json:"name"`

	// Description explains what this category represents.
	Description string `json:"description,omitempty"`

	// Color is used for visual grouping (hex or named color).
	Color string `json:"color,omitempty"`

	// Importance is the static weight for this category (critical, high, medium, low).
	// Used for "-ilities" prioritization (security, availability, resiliency, etc.).
	Importance string `json:"importance,omitempty"`
}
