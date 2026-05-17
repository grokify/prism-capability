package capstack

// FrameworkMapping maps a capability to compliance/security framework controls.
type FrameworkMapping struct {
	// Framework is the framework name (nist-csf-2.0, iso-27001, etc.).
	Framework string `json:"framework"`

	// Controls lists the control IDs from the framework.
	Controls []string `json:"controls"`
}
