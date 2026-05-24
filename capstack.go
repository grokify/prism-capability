package capstack

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
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

	// MarketIntegration configures global market-strategy-engine integration.
	MarketIntegration *MarketIntegration `json:"marketIntegration,omitempty"`

	// PriorityFramework defines the priority tier structure for organizing capabilities.
	// When present, an HTML page explaining the priority tiers can be generated.
	PriorityFramework *PriorityFramework `json:"priorityFramework,omitempty"`
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

// CapabilitiesByMarketCapability returns org capabilities that enable a market capability.
func (cs *CapabilityStack) CapabilitiesByMarketCapability(marketCapID string) []Capability {
	var result []Capability
	for _, cap := range cs.AllCapabilities() {
		if cap.MarketRef != nil {
			for _, mcid := range cap.MarketRef.CapabilityIDs {
				if mcid == marketCapID {
					result = append(result, cap)
					break
				}
			}
		}
	}
	return result
}

// CapabilitiesByMarket returns org capabilities linked to a specific market.
func (cs *CapabilityStack) CapabilitiesByMarket(marketID string) []Capability {
	var result []Capability
	for _, cap := range cs.AllCapabilities() {
		if cap.MarketRef != nil && cap.MarketRef.MarketID == marketID {
			result = append(result, cap)
		}
	}
	return result
}

// CapabilitiesForSegment returns org capabilities that benefit a market segment.
func (cs *CapabilityStack) CapabilitiesForSegment(segmentID string) []Capability {
	var result []Capability
	for _, cap := range cs.AllCapabilities() {
		if cap.MarketRef != nil {
			for _, seg := range cap.MarketRef.Segments {
				if seg == segmentID {
					result = append(result, cap)
					break
				}
			}
		}
	}
	return result
}

// SortMethod defines how capabilities should be sorted.
type SortMethod string

const (
	// SortByOrder sorts by explicit Order field (default).
	SortByOrder SortMethod = "order"
	// SortByName sorts alphabetically by Name.
	SortByName SortMethod = "name"
	// SortByImportance sorts by Importance weight (critical first).
	SortByImportance SortMethod = "importance"
	// SortByPriority sorts by Priority (critical first, for static priority field).
	SortByPriority SortMethod = "priority"
	// SortByStatus sorts by Status (operational first).
	SortByStatus SortMethod = "status"
)

// AllSortMethods returns all valid sort methods.
func AllSortMethods() []SortMethod {
	return []SortMethod{
		SortByOrder,
		SortByName,
		SortByImportance,
		SortByPriority,
		SortByStatus,
	}
}

// ValidSortMethod checks if a sort method is valid.
func ValidSortMethod(method string) bool {
	for _, m := range AllSortMethods() {
		if string(m) == method {
			return true
		}
	}
	return false
}

// SortCapabilities sorts the capabilities slice in place by the given method.
func (cs *CapabilityStack) SortCapabilities(method SortMethod) {
	sort.SliceStable(cs.Capabilities, func(i, j int) bool {
		return compareCapabilities(cs.Capabilities[i], cs.Capabilities[j], method)
	})
	sort.SliceStable(cs.Foundational, func(i, j int) bool {
		return compareCapabilities(cs.Foundational[i], cs.Foundational[j], method)
	})
}

// SortedCapabilities returns a sorted copy of all capabilities.
func (cs *CapabilityStack) SortedCapabilities(method SortMethod) []Capability {
	caps := cs.AllCapabilities()
	sort.SliceStable(caps, func(i, j int) bool {
		return compareCapabilities(caps[i], caps[j], method)
	})
	return caps
}

// compareCapabilities returns true if a should come before b.
func compareCapabilities(a, b Capability, method SortMethod) bool {
	switch method {
	case SortByOrder:
		// If both have explicit order, use it
		if a.Order != 0 && b.Order != 0 {
			return a.Order < b.Order
		}
		// Items with explicit order come before items without
		if a.Order != 0 {
			return true
		}
		if b.Order != 0 {
			return false
		}
		// Fall back to name for items without order
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)

	case SortByName:
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)

	case SortByImportance:
		// Higher importance first (critical=4, high=3, medium=2, low=1)
		aWeight := ImportanceWeight(a.Importance)
		bWeight := ImportanceWeight(b.Importance)
		if aWeight != bWeight {
			return aWeight > bWeight
		}
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)

	case SortByPriority:
		// Higher priority first (critical=4, high=3, medium=2, low=1)
		aWeight := PriorityWeight(a.Priority)
		bWeight := PriorityWeight(b.Priority)
		if aWeight != bWeight {
			return aWeight > bWeight
		}
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)

	case SortByStatus:
		// Operational first, then implemented, in-progress, planned, deprecated
		aWeight := statusWeight(a.Status)
		bWeight := statusWeight(b.Status)
		if aWeight != bWeight {
			return aWeight > bWeight
		}
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)

	default:
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)
	}
}

// statusWeight returns a sort weight for capability status.
func statusWeight(status string) int {
	switch status {
	case StatusOperational:
		return 5
	case StatusImplemented:
		return 4
	case StatusInProgress:
		return 3
	case StatusPlanned:
		return 2
	case StatusDeprecated:
		return 1
	default:
		return 0
	}
}
