package render

import (
	"strings"
	"testing"

	capstack "github.com/plexusone/capability-stack-spec"
)

func TestRenderHTML(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{
			Name:    "test-stack",
			Version: "1.0.0",
			Title:   "Test Stack",
		},
		Layers: []capstack.Layer{
			{ID: "layer-1", Name: "Layer One"},
		},
		Capabilities: []capstack.Capability{
			{
				ID:      "cap-a",
				Name:    "Capability A",
				LayerID: "layer-1",
				Status:  capstack.StatusOperational,
				Owner:   "Team A",
			},
		},
	}

	opts := DefaultHTMLOptions()
	result, err := RenderHTMLString(doc, opts)
	if err != nil {
		t.Fatalf("RenderHTMLString() error = %v", err)
	}

	// Should be a fragment (no DOCTYPE)
	if strings.Contains(result, "<!DOCTYPE") {
		t.Error("Default HTML should be a fragment without DOCTYPE")
	}

	// Verify key elements
	checks := []string{
		"cs-container",
		"Test Stack",
		"Layer One",
		"Capability A",
		"#10b981", // Operational color
		"cs-legend",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("Output missing expected content: %q", check)
		}
	}
}

func TestRenderHTMLStandalone(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{
			Name:    "test",
			Version: "1.0.0",
			Title:   "Standalone Test",
		},
		Layers: []capstack.Layer{
			{ID: "layer-1", Name: "Layer"},
		},
		Capabilities: []capstack.Capability{
			{ID: "cap-a", Name: "Cap A", LayerID: "layer-1"},
		},
	}

	opts := DefaultHTMLOptions()
	opts.Standalone = true

	result, err := RenderHTMLString(doc, opts)
	if err != nil {
		t.Fatalf("RenderHTMLString() error = %v", err)
	}

	// Should have full HTML document structure
	checks := []string{
		"<!DOCTYPE html>",
		"<html lang=\"en\">",
		"<head>",
		"<meta charset=\"UTF-8\">",
		"<title>Standalone Test</title>",
		"<style>",
		"</head>",
		"<body",
		"</body>",
		"</html>",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("Standalone HTML missing: %q", check)
		}
	}
}

func TestRenderHTMLDarkTheme(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "test", Version: "1.0.0"},
		Layers:   []capstack.Layer{{ID: "layer-1", Name: "Layer"}},
		Capabilities: []capstack.Capability{
			{ID: "cap-a", Name: "Cap A", LayerID: "layer-1"},
		},
	}

	opts := DefaultHTMLOptions()
	opts.DarkTheme = true
	opts.Standalone = true

	result, err := RenderHTMLString(doc, opts)
	if err != nil {
		t.Fatalf("RenderHTMLString() error = %v", err)
	}

	// Should have dark class
	if !strings.Contains(result, "class=\"dark\"") && !strings.Contains(result, "cs-container dark") {
		t.Error("Dark theme HTML should have dark class")
	}
}

func TestRenderHTMLNoLegend(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "test", Version: "1.0.0"},
		Layers:   []capstack.Layer{{ID: "layer-1", Name: "Layer"}},
		Capabilities: []capstack.Capability{
			{ID: "cap-a", Name: "Cap A", LayerID: "layer-1"},
		},
	}

	opts := DefaultHTMLOptions()
	opts.ShowLegend = false

	result, err := RenderHTMLString(doc, opts)
	if err != nil {
		t.Fatalf("RenderHTMLString() error = %v", err)
	}

	if strings.Contains(result, "cs-legend") {
		t.Error("Should not contain legend when ShowLegend=false")
	}
}

//nolint:dupl // Test code duplication is acceptable for clarity
func TestRenderHTMLFoundational(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "test", Version: "1.0.0"},
		Layers:   []capstack.Layer{{ID: "layer-1", Name: "Layer"}},
		Capabilities: []capstack.Capability{
			{ID: "cap-a", Name: "Cap A", LayerID: "layer-1"},
		},
		Foundational: []capstack.Capability{
			{ID: "found-1", Name: "Foundation", LayerID: "layer-1", Status: capstack.StatusOperational},
		},
	}

	// With foundational
	opts := DefaultHTMLOptions()
	result, err := RenderHTMLString(doc, opts)
	if err != nil {
		t.Fatalf("RenderHTMLString() error = %v", err)
	}

	if !strings.Contains(result, "cs-foundational") {
		t.Error("Expected foundational section in output")
	}
	if !strings.Contains(result, "Foundation") {
		t.Error("Expected foundational capability in output")
	}

	// Without foundational
	opts.ShowFoundational = false
	result, err = RenderHTMLString(doc, opts)
	if err != nil {
		t.Fatalf("RenderHTMLString() error = %v", err)
	}

	if strings.Contains(result, "cs-foundational") {
		t.Error("Expected no foundational section when ShowFoundational=false")
	}
}

func TestRenderHTMLTooltip(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "test", Version: "1.0.0"},
		Layers:   []capstack.Layer{{ID: "layer-1", Name: "Layer"}},
		Capabilities: []capstack.Capability{
			{
				ID:          "cap-a",
				Name:        "Cap A",
				FullName:    "Capability Alpha",
				Description: "Test description",
				LayerID:     "layer-1",
				Status:      capstack.StatusImplemented,
				Owner:       "Team Alpha",
			},
		},
	}

	opts := DefaultHTMLOptions()
	result, err := RenderHTMLString(doc, opts)
	if err != nil {
		t.Fatalf("RenderHTMLString() error = %v", err)
	}

	// Tooltip should contain capability details
	if !strings.Contains(result, "Capability Alpha") {
		t.Error("Tooltip should contain full name")
	}
	if !strings.Contains(result, "Test description") {
		t.Error("Tooltip should contain description")
	}
	if !strings.Contains(result, "Owner: Team Alpha") {
		t.Error("Tooltip should contain owner")
	}
	if !strings.Contains(result, "Status: implemented") {
		t.Error("Tooltip should contain status")
	}
}

func TestDefaultHTMLOptions(t *testing.T) {
	opts := DefaultHTMLOptions()

	if !opts.ShowLegend {
		t.Error("DefaultHTMLOptions().ShowLegend should be true")
	}
	if !opts.ShowFoundational {
		t.Error("DefaultHTMLOptions().ShowFoundational should be true")
	}
	if opts.Standalone {
		t.Error("DefaultHTMLOptions().Standalone should be false")
	}
	if opts.DarkTheme {
		t.Error("DefaultHTMLOptions().DarkTheme should be false")
	}
}
