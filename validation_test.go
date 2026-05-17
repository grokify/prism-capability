package capstack

import (
	"testing"
)

func TestValidateKebabCase(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"valid-id", true},
		{"also-valid-123", true},
		{"single", true},
		{"a1-b2-c3", true},
		{"Invalid", false},
		{"invalid_underscore", false},
		{"invalid space", false},
		{"-invalid", false},
		{"invalid-", false},
		{"123-invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ValidateKebabCase(tt.input)
			if got != tt.want {
				t.Errorf("ValidateKebabCase(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateCapabilityStatus(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"planned", false},
		{"in-progress", false},
		{"implemented", false},
		{"operational", false},
		{"deprecated", false},
		{"", false}, // Optional field
		{"invalid", true},
		{"PLANNED", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			err := ValidateCapabilityStatus(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCapabilityStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestCapabilityValidate(t *testing.T) {
	tests := []struct {
		name    string
		cap     Capability
		wantErr bool
	}{
		{
			name: "valid capability",
			cap: Capability{
				ID:      "test-cap",
				Name:    "Test Capability",
				LayerID: "test-layer",
			},
			wantErr: false,
		},
		{
			name: "missing id",
			cap: Capability{
				Name:    "Test",
				LayerID: "test-layer",
			},
			wantErr: true,
		},
		{
			name: "invalid id format",
			cap: Capability{
				ID:      "InvalidID",
				Name:    "Test",
				LayerID: "test-layer",
			},
			wantErr: true,
		},
		{
			name: "missing name",
			cap: Capability{
				ID:      "test-cap",
				LayerID: "test-layer",
			},
			wantErr: true,
		},
		{
			name: "missing layerId",
			cap: Capability{
				ID:   "test-cap",
				Name: "Test",
			},
			wantErr: true,
		},
		{
			name: "invalid status",
			cap: Capability{
				ID:      "test-cap",
				Name:    "Test",
				LayerID: "test-layer",
				Status:  "invalid-status",
			},
			wantErr: true,
		},
		{
			name: "invalid priority",
			cap: Capability{
				ID:       "test-cap",
				Name:     "Test",
				LayerID:  "test-layer",
				Priority: "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid tag format",
			cap: Capability{
				ID:      "test-cap",
				Name:    "Test",
				LayerID: "test-layer",
				Tags:    []string{"valid-tag", "Invalid_Tag"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.cap.Validate()
			if errs.HasErrors() != tt.wantErr {
				t.Errorf("Capability.Validate() errors = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestCapabilityStackValidate(t *testing.T) {
	validStack := CapabilityStack{
		Metadata: Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
		},
		Layers: []Layer{
			{ID: "layer-1", Name: "Layer 1"},
		},
		Categories: []Category{
			{ID: "cat-1", Name: "Category 1"},
		},
		Capabilities: []Capability{
			{ID: "cap-1", Name: "Cap 1", LayerID: "layer-1", CategoryID: "cat-1"},
		},
	}

	errs := validStack.Validate()
	if errs.HasErrors() {
		t.Errorf("Valid stack should not have errors: %v", errs)
	}
}

func TestCapabilityStackValidateDuplicateIDs(t *testing.T) {
	stack := CapabilityStack{
		Metadata: Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
		},
		Layers: []Layer{
			{ID: "layer-1", Name: "Layer 1"},
			{ID: "layer-1", Name: "Layer 1 Duplicate"},
		},
		Capabilities: []Capability{
			{ID: "cap-1", Name: "Cap 1", LayerID: "layer-1"},
		},
	}

	errs := stack.Validate()
	if !errs.HasErrors() {
		t.Error("Expected duplicate layer ID error")
	}

	found := false
	for _, e := range errs {
		if e.Message == "duplicate ID, also used at layers[0]" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected duplicate ID error, got: %v", errs)
	}
}

func TestCapabilityStackValidateInvalidReferences(t *testing.T) {
	stack := CapabilityStack{
		Metadata: Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
		},
		Layers: []Layer{
			{ID: "layer-1", Name: "Layer 1"},
		},
		Capabilities: []Capability{
			{ID: "cap-1", Name: "Cap 1", LayerID: "non-existent-layer"},
		},
	}

	errs := stack.Validate()
	if !errs.HasErrors() {
		t.Error("Expected invalid layer reference error")
	}

	found := false
	for _, e := range errs {
		if e.Message == "references non-existent layer ID" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected layer reference error, got: %v", errs)
	}
}

func TestCapabilityStackValidateDependencyCycle(t *testing.T) {
	stack := CapabilityStack{
		Metadata: Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
		},
		Layers: []Layer{
			{ID: "layer-1", Name: "Layer 1"},
		},
		Capabilities: []Capability{
			{ID: "cap-a", Name: "Cap A", LayerID: "layer-1", Dependencies: []string{"cap-b"}},
			{ID: "cap-b", Name: "Cap B", LayerID: "layer-1", Dependencies: []string{"cap-c"}},
			{ID: "cap-c", Name: "Cap C", LayerID: "layer-1", Dependencies: []string{"cap-a"}},
		},
	}

	errs := stack.Validate()
	if !errs.HasErrors() {
		t.Error("Expected dependency cycle error")
	}

	found := false
	for _, e := range errs {
		if e.Message == "circular dependency detected" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected cycle detection error, got: %v", errs)
	}
}

func TestValidationErrorsError(t *testing.T) {
	errs := ValidationErrors{
		{Field: "name", Message: "is required"},
		{Field: "id", Value: "bad", Message: "must be kebab-case"},
	}

	got := errs.Error()
	want := `name: is required; id: must be kebab-case (value: "bad")`
	if got != want {
		t.Errorf("ValidationErrors.Error() = %q, want %q", got, want)
	}
}
