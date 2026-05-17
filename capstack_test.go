package capstack

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	// Test loading the operations example
	doc, err := LoadFromFile("examples/operations-saas-greenfield.json")
	if err != nil {
		t.Fatalf("LoadFromFile() error = %v", err)
	}

	if doc.Metadata.Name != "b2b-saas-operations" {
		t.Errorf("Metadata.Name = %q, want %q", doc.Metadata.Name, "b2b-saas-operations")
	}

	if doc.Metadata.Domain != "operations" {
		t.Errorf("Metadata.Domain = %q, want %q", doc.Metadata.Domain, "operations")
	}

	if len(doc.Layers) == 0 {
		t.Error("Expected at least one layer")
	}

	if len(doc.Capabilities) == 0 {
		t.Error("Expected at least one capability")
	}
}

func TestLoadFromFileValidation(t *testing.T) {
	doc, err := LoadFromFile("examples/operations-saas-greenfield.json")
	if err != nil {
		t.Fatalf("LoadFromFile() error = %v", err)
	}

	errs := doc.Validate()
	if errs.HasErrors() {
		t.Errorf("Example file should be valid, got errors: %v", errs)
	}
}

func TestGetLayerByID(t *testing.T) {
	doc := &CapabilityStack{
		Layers: []Layer{
			{ID: "layer-1", Name: "Layer 1"},
			{ID: "layer-2", Name: "Layer 2"},
		},
	}

	layer := doc.GetLayerByID("layer-1")
	if layer == nil {
		t.Fatal("Expected to find layer-1")
	}
	if layer.Name != "Layer 1" {
		t.Errorf("Layer.Name = %q, want %q", layer.Name, "Layer 1")
	}

	notFound := doc.GetLayerByID("non-existent")
	if notFound != nil {
		t.Error("Expected nil for non-existent layer")
	}
}

func TestGetCapabilityByID(t *testing.T) {
	doc := &CapabilityStack{
		Capabilities: []Capability{
			{ID: "cap-1", Name: "Cap 1"},
		},
		Foundational: []Capability{
			{ID: "found-1", Name: "Foundational 1"},
		},
	}

	// Find in capabilities
	cap := doc.GetCapabilityByID("cap-1")
	if cap == nil {
		t.Fatal("Expected to find cap-1")
	}

	// Find in foundational
	found := doc.GetCapabilityByID("found-1")
	if found == nil {
		t.Fatal("Expected to find found-1")
	}

	// Not found
	notFound := doc.GetCapabilityByID("non-existent")
	if notFound != nil {
		t.Error("Expected nil for non-existent capability")
	}
}

func TestCapabilitiesByStatus(t *testing.T) {
	doc := &CapabilityStack{
		Capabilities: []Capability{
			{ID: "cap-1", Name: "Cap 1", Status: StatusPlanned},
			{ID: "cap-2", Name: "Cap 2", Status: StatusImplemented},
			{ID: "cap-3", Name: "Cap 3", Status: StatusPlanned},
		},
	}

	planned := doc.CapabilitiesByStatus(StatusPlanned)
	if len(planned) != 2 {
		t.Errorf("Expected 2 planned capabilities, got %d", len(planned))
	}

	implemented := doc.CapabilitiesByStatus(StatusImplemented)
	if len(implemented) != 1 {
		t.Errorf("Expected 1 implemented capability, got %d", len(implemented))
	}
}

func TestCapabilitiesByTag(t *testing.T) {
	doc := &CapabilityStack{
		Capabilities: []Capability{
			{ID: "cap-1", Name: "Cap 1", Tags: []string{"shift-left", "security"}},
			{ID: "cap-2", Name: "Cap 2", Tags: []string{"automation"}},
			{ID: "cap-3", Name: "Cap 3", Tags: []string{"shift-left"}},
		},
	}

	shiftLeft := doc.CapabilitiesByTag("shift-left")
	if len(shiftLeft) != 2 {
		t.Errorf("Expected 2 shift-left capabilities, got %d", len(shiftLeft))
	}

	automation := doc.CapabilitiesByTag("automation")
	if len(automation) != 1 {
		t.Errorf("Expected 1 automation capability, got %d", len(automation))
	}

	none := doc.CapabilitiesByTag("non-existent")
	if len(none) != 0 {
		t.Errorf("Expected 0 capabilities, got %d", len(none))
	}
}

func TestAllCapabilities(t *testing.T) {
	doc := &CapabilityStack{
		Capabilities: []Capability{
			{ID: "cap-1"},
			{ID: "cap-2"},
		},
		Foundational: []Capability{
			{ID: "found-1"},
		},
	}

	all := doc.AllCapabilities()
	if len(all) != 3 {
		t.Errorf("Expected 3 capabilities, got %d", len(all))
	}
}

func TestSaveToFile(t *testing.T) {
	doc := &CapabilityStack{
		Metadata: Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
		},
		Layers: []Layer{
			{ID: "layer-1", Name: "Layer 1"},
		},
		Capabilities: []Capability{
			{ID: "cap-1", Name: "Cap 1", LayerID: "layer-1"},
		},
	}

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.json")

	err := doc.SaveToFile(path)
	if err != nil {
		t.Fatalf("SaveToFile() error = %v", err)
	}

	// Read it back
	loaded, err := LoadFromFile(path)
	if err != nil {
		t.Fatalf("LoadFromFile() error = %v", err)
	}

	if loaded.Metadata.Name != doc.Metadata.Name {
		t.Errorf("Loaded name = %q, want %q", loaded.Metadata.Name, doc.Metadata.Name)
	}
}

func TestJSONRoundTrip(t *testing.T) {
	original := &CapabilityStack{
		Schema: "https://example.com/schema.json",
		Metadata: Metadata{
			Name:        "test",
			Version:     "1.0.0",
			Title:       "Test Stack",
			Description: "A test",
			Domain:      "operations",
			Authors:     []string{"Author 1"},
		},
		Layers: []Layer{
			{ID: "layer-1", Name: "Layer 1", Phase: "build"},
		},
		Categories: []Category{
			{ID: "cat-1", Name: "Category 1", Color: "#ff0000"},
		},
		Capabilities: []Capability{
			{
				ID:         "cap-1",
				Name:       "Cap 1",
				FullName:   "Capability One",
				LayerID:    "layer-1",
				CategoryID: "cat-1",
				Status:     StatusImplemented,
				Priority:   PriorityHigh,
				Owner:      "Team",
				Tags:       []string{"tag-1"},
				Tooling:    []Tool{{Name: "Tool 1", Type: ToolTypeOpenSource}},
				PRISMRef:   &PRISMRef{DomainID: "operations", SLIIDs: []string{"sli-1"}},
			},
		},
		PRISMIntegration: &PRISMIntegration{
			ModelRef:      "./model.json",
			DefaultDomain: "operations",
		},
	}

	data, err := json.MarshalIndent(original, "", "  ")
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}

	var loaded CapabilityStack
	if err := json.Unmarshal(data, &loaded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	// Verify key fields
	if loaded.Metadata.Name != original.Metadata.Name {
		t.Errorf("Name mismatch: got %q, want %q", loaded.Metadata.Name, original.Metadata.Name)
	}
	if len(loaded.Capabilities) != len(original.Capabilities) {
		t.Errorf("Capability count mismatch: got %d, want %d", len(loaded.Capabilities), len(original.Capabilities))
	}
	if loaded.PRISMIntegration == nil {
		t.Error("PRISMIntegration should not be nil")
	}
}

func TestLoadSecurityExample(t *testing.T) {
	// Test loading the security example if it exists
	path := "examples/security-saas-greenfield.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Skip("Security example not found")
	}

	doc, err := LoadFromFile(path)
	if err != nil {
		t.Fatalf("LoadFromFile() error = %v", err)
	}

	errs := doc.Validate()
	if errs.HasErrors() {
		t.Errorf("Security example should be valid, got errors: %v", errs)
	}
}
