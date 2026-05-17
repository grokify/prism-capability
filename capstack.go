package capstack

import (
	"encoding/json"
	"os"
)

// CapabilityStack is the root document for a capability stack specification.
type CapabilityStack struct {
	// Schema is the JSON Schema reference.
	Schema string `json:"$schema,omitempty"`

	// Metadata contains document-level information.
	Metadata Metadata `json:"metadata"`

	// Layers are ordered list of layers (rows) in the capability stack.
	Layers []Layer `json:"layers"`

	// Categories define groupings for capabilities within layers.
	Categories []Category `json:"categories,omitempty"`

	// Capabilities are all capabilities in the stack.
	Capabilities []Capability `json:"capabilities"`

	// Foundational are cross-cutting capabilities that span multiple layers.
	Foundational []Capability `json:"foundational,omitempty"`

	// PRISMIntegration configures global PRISM integration.
	PRISMIntegration *PRISMIntegration `json:"prismIntegration,omitempty"`
}

// LoadFromFile reads a CapabilityStack from a JSON file.
func LoadFromFile(path string) (*CapabilityStack, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc CapabilityStack
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

// SaveToFile writes the CapabilityStack to a JSON file.
func (cs *CapabilityStack) SaveToFile(path string) error {
	data, err := json.MarshalIndent(cs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

// GetLayerByID returns a layer by its ID, or nil if not found.
func (cs *CapabilityStack) GetLayerByID(id string) *Layer {
	for i := range cs.Layers {
		if cs.Layers[i].ID == id {
			return &cs.Layers[i]
		}
	}
	return nil
}

// GetCategoryByID returns a category by its ID, or nil if not found.
func (cs *CapabilityStack) GetCategoryByID(id string) *Category {
	for i := range cs.Categories {
		if cs.Categories[i].ID == id {
			return &cs.Categories[i]
		}
	}
	return nil
}

// GetCapabilityByID returns a capability by its ID, or nil if not found.
// Searches both capabilities and foundational lists.
func (cs *CapabilityStack) GetCapabilityByID(id string) *Capability {
	for i := range cs.Capabilities {
		if cs.Capabilities[i].ID == id {
			return &cs.Capabilities[i]
		}
	}
	for i := range cs.Foundational {
		if cs.Foundational[i].ID == id {
			return &cs.Foundational[i]
		}
	}
	return nil
}

// AllCapabilities returns all capabilities including foundational ones.
func (cs *CapabilityStack) AllCapabilities() []Capability {
	all := make([]Capability, 0, len(cs.Capabilities)+len(cs.Foundational))
	all = append(all, cs.Capabilities...)
	all = append(all, cs.Foundational...)
	return all
}

// CapabilitiesByLayer returns capabilities belonging to a specific layer.
func (cs *CapabilityStack) CapabilitiesByLayer(layerID string) []Capability {
	var result []Capability
	for _, cap := range cs.AllCapabilities() {
		if cap.LayerID == layerID {
			result = append(result, cap)
		}
	}
	return result
}

// CapabilitiesByCategory returns capabilities belonging to a specific category.
func (cs *CapabilityStack) CapabilitiesByCategory(categoryID string) []Capability {
	var result []Capability
	for _, cap := range cs.AllCapabilities() {
		if cap.CategoryID == categoryID {
			result = append(result, cap)
		}
	}
	return result
}

// CapabilitiesByStatus returns capabilities with a specific status.
func (cs *CapabilityStack) CapabilitiesByStatus(status string) []Capability {
	var result []Capability
	for _, cap := range cs.AllCapabilities() {
		if cap.Status == status {
			result = append(result, cap)
		}
	}
	return result
}

// CapabilitiesByTag returns capabilities with a specific tag.
func (cs *CapabilityStack) CapabilitiesByTag(tag string) []Capability {
	var result []Capability
	for _, cap := range cs.AllCapabilities() {
		for _, t := range cap.Tags {
			if t == tag {
				result = append(result, cap)
				break
			}
		}
	}
	return result
}

// LayerIDs returns all layer IDs in order.
func (cs *CapabilityStack) LayerIDs() []string {
	ids := make([]string, len(cs.Layers))
	for i, layer := range cs.Layers {
		ids[i] = layer.ID
	}
	return ids
}

// CategoryIDs returns all category IDs.
func (cs *CapabilityStack) CategoryIDs() []string {
	ids := make([]string, len(cs.Categories))
	for i, cat := range cs.Categories {
		ids[i] = cat.ID
	}
	return ids
}

// CapabilityIDs returns all capability IDs including foundational.
func (cs *CapabilityStack) CapabilityIDs() []string {
	all := cs.AllCapabilities()
	ids := make([]string, len(all))
	for i, cap := range all {
		ids[i] = cap.ID
	}
	return ids
}
