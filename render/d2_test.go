package render

import (
	"strings"
	"testing"

	capstack "github.com/plexusone/capability-stack-spec"
)

func TestRenderD2(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
			Title:   "Test Stack",
		},
		Layers: []capstack.Layer{
			{ID: "layer-1", Name: "Layer One"},
			{ID: "layer-2", Name: "Layer Two"},
		},
		Categories: []capstack.Category{
			{ID: "cat-1", Name: "Category One", Color: "#ff0000"},
		},
		Capabilities: []capstack.Capability{
			{
				ID:         "cap-a",
				Name:       "Capability A",
				LayerID:    "layer-1",
				CategoryID: "cat-1",
				Status:     capstack.StatusOperational,
			},
			{
				ID:           "cap-b",
				Name:         "Capability B",
				LayerID:      "layer-2",
				Status:       capstack.StatusPlanned,
				Dependencies: []string{"cap-a"},
			},
		},
	}

	opts := DefaultD2Options()
	result, err := RenderD2String(doc, opts)
	if err != nil {
		t.Fatalf("RenderD2String() error = %v", err)
	}

	// Verify key elements are present
	checks := []string{
		"Test Stack",                     // Title
		"layer_1: Layer One",             // Layer container
		"layer_2: Layer Two",             // Layer container
		"cap_a: Capability A",            // Capability
		"cap_b: Capability B",            // Capability
		"#10b981",                        // Operational color (green)
		"#9ca3af",                        // Planned color (gray)
		"layer_1.cap_a <- layer_2.cap_b", // Dependency arrow
		"legend:",                        // Legend
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("Output missing expected content: %q", check)
		}
	}
}

func TestRenderD2ColorByCategory(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{
			Name:    "test",
			Version: "1.0.0",
		},
		Layers: []capstack.Layer{
			{ID: "layer-1", Name: "Layer One"},
		},
		Categories: []capstack.Category{
			{ID: "cat-1", Name: "Category", Color: "#abcdef"},
		},
		Capabilities: []capstack.Capability{
			{
				ID:         "cap-a",
				Name:       "Cap A",
				LayerID:    "layer-1",
				CategoryID: "cat-1",
				Status:     capstack.StatusOperational,
			},
		},
	}

	opts := DefaultD2Options()
	opts.ColorByStatus = false

	result, err := RenderD2String(doc, opts)
	if err != nil {
		t.Fatalf("RenderD2String() error = %v", err)
	}

	// Should use category color instead of status color
	if !strings.Contains(result, "#abcdef") {
		t.Error("Expected category color #abcdef in output")
	}
}

func TestRenderD2NoDependencies(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "test", Version: "1.0.0"},
		Layers:   []capstack.Layer{{ID: "layer-1", Name: "Layer"}},
		Capabilities: []capstack.Capability{
			{ID: "cap-a", Name: "A", LayerID: "layer-1"},
			{ID: "cap-b", Name: "B", LayerID: "layer-1", Dependencies: []string{"cap-a"}},
		},
	}

	opts := DefaultD2Options()
	opts.ShowDependencies = false

	result, err := RenderD2String(doc, opts)
	if err != nil {
		t.Fatalf("RenderD2String() error = %v", err)
	}

	// Should not contain dependency arrow
	if strings.Contains(result, "<-") {
		t.Error("Expected no dependency arrows when ShowDependencies=false")
	}
}

//nolint:dupl // Test code duplication is acceptable for clarity
func TestRenderD2Foundational(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "test", Version: "1.0.0"},
		Layers:   []capstack.Layer{{ID: "layer-1", Name: "Layer"}},
		Capabilities: []capstack.Capability{
			{ID: "cap-a", Name: "A", LayerID: "layer-1"},
		},
		Foundational: []capstack.Capability{
			{ID: "found-1", Name: "Foundation", LayerID: "layer-1", Status: capstack.StatusOperational},
		},
	}

	// With foundational
	opts := DefaultD2Options()
	result, err := RenderD2String(doc, opts)
	if err != nil {
		t.Fatalf("RenderD2String() error = %v", err)
	}

	if !strings.Contains(result, "foundational: Foundational") {
		t.Error("Expected foundational section in output")
	}
	if !strings.Contains(result, "found_1: Foundation") {
		t.Error("Expected foundational capability in output")
	}

	// Without foundational
	opts.ShowFoundational = false
	result, err = RenderD2String(doc, opts)
	if err != nil {
		t.Fatalf("RenderD2String() error = %v", err)
	}

	if strings.Contains(result, "foundational:") {
		t.Error("Expected no foundational section when ShowFoundational=false")
	}
}

func TestSanitizeD2ID(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"simple", "simple"},
		{"kebab-case", "kebab_case"},
		{"multi-part-id", "multi_part_id"},
	}

	for _, tt := range tests {
		got := sanitizeD2ID(tt.input)
		if got != tt.want {
			t.Errorf("sanitizeD2ID(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestRenderD2Grid(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
			Title:   "Executive View",
		},
		Layers: []capstack.Layer{
			{ID: "layer-1", Name: "Layer One"},
			{ID: "layer-2", Name: "Layer Two"},
		},
		Capabilities: []capstack.Capability{
			{ID: "cap-a", Name: "Cap A", LayerID: "layer-1", Status: capstack.StatusOperational},
			{ID: "cap-b", Name: "Cap B", LayerID: "layer-1", Status: capstack.StatusImplemented},
			{ID: "cap-c", Name: "Cap C", LayerID: "layer-2", Status: capstack.StatusPlanned},
		},
	}

	opts := GridD2Options()
	result, err := RenderD2String(doc, opts)
	if err != nil {
		t.Fatalf("RenderD2String() error = %v", err)
	}

	// Verify grid-specific elements
	checks := []string{
		"Executive View",  // Title
		"stack:",          // Outer stacking container
		"grid-rows: 2",    // Vertical stacking (2 layers)
		"grid-columns: 6", // Grid layout within layer
		"style.shadow",    // Shadow for clean look
		"layer_1:",        // Layer container
		"cap_a: Cap A",    // Capability
		"bottom-center",   // Compact legend position
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("Grid output missing expected content: %q", check)
		}
	}

	// Grid view should NOT have dependency arrows
	if strings.Contains(result, "<-") {
		t.Error("Grid view should not contain dependency arrows")
	}
}

func TestGridD2Options(t *testing.T) {
	opts := GridD2Options()

	if opts.Style != D2StyleGrid {
		t.Errorf("GridD2Options().Style = %v, want %v", opts.Style, D2StyleGrid)
	}
	if opts.ShowDependencies {
		t.Error("GridD2Options().ShowDependencies should be false")
	}
	if !opts.ShowLegend {
		t.Error("GridD2Options().ShowLegend should be true")
	}
}
