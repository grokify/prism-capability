package capstack

// Layer represents a row in the capability stack, typically mapping to
// a lifecycle phase or organizational boundary.
type Layer struct {
	// ID is the unique identifier for the layer (kebab-case).
	ID string `json:"id"`

	// Name is the display name for the layer.
	Name string `json:"name"`

	// Description explains the purpose/objective of this layer.
	Description string `json:"description,omitempty"`

	// Order is the sort order (1 = top layer).
	Order int `json:"order,omitempty"`

	// Phase is the SDLC or lifecycle phase this layer represents.
	Phase string `json:"phase,omitempty"`

	// NistCsfFunction maps the layer to a NIST CSF 2.0 function.
	NistCsfFunction string `json:"nistCsfFunction,omitempty"`
}
