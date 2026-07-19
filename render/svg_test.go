package render

import (
	"strings"
	"testing"

	capstack "github.com/grokify/prism-capability"
)

func svgTestStack() *capstack.CapabilityStack {
	return &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "test-platform", Version: "1.0.0", Title: "Test Platform"},
		Layers: []capstack.Layer{
			{ID: "experiences", Name: "Experiences", Order: 1},
			{ID: "intelligence", Name: "Intelligence", Order: 2},
			{ID: "trust", Name: "Trust", Order: 3},
		},
		Capabilities: []capstack.Capability{
			{ID: "chat", Name: "Chat", LayerID: "experiences", Tooling: []capstack.Tool{{Name: "OmniChat"}}},
			{ID: "voice", Name: "Voice", LayerID: "experiences", Tooling: []capstack.Tool{{Name: "OmniVoice"}}},
			{ID: "llm", Name: "LLM Abstraction", LayerID: "intelligence", Tooling: []capstack.Tool{{Name: "OmniLLM"}}},
			{ID: "secrets", Name: "Secret Management", LayerID: "trust", Tooling: []capstack.Tool{{Name: "OmniVault"}}},
		},
	}
}

func TestRenderSVGStack(t *testing.T) {
	doc := svgTestStack()
	opts := DefaultSVGOptions()
	opts.Substrate = "Go · providers"

	out, err := RenderSVGString(doc, opts)
	if err != nil {
		t.Fatalf("RenderSVGString: %v", err)
	}
	for _, want := range []string{"<svg", "viewBox", "Test Platform", "Experiences", "Secret Management", "Go · providers", "</svg>"} {
		if !strings.Contains(out, want) {
			t.Errorf("stack SVG missing %q", want)
		}
	}
	if !strings.HasSuffix(strings.TrimSpace(out), "</svg>") {
		t.Errorf("stack SVG not well-terminated")
	}
}

func TestRenderSVGHub(t *testing.T) {
	doc := svgTestStack()
	opts := DefaultSVGOptions()
	opts.Layout = SVGLayoutHub
	opts.CenterLabel = "MyAgent"
	opts.CenterSubLabel = "Runtime"

	out, err := RenderSVGString(doc, opts)
	if err != nil {
		t.Fatalf("RenderSVGString: %v", err)
	}
	// Hub shows layer names, the center label, and representative tooling.
	for _, want := range []string{"<svg", "polygon", "MyAgent", "Runtime", "Experiences", "OmniChat", "</svg>"} {
		if !strings.Contains(out, want) {
			t.Errorf("hub SVG missing %q", want)
		}
	}
}

func TestRenderSVGHubFallsBackToCapabilityCount(t *testing.T) {
	doc := &capstack.CapabilityStack{
		Metadata: capstack.Metadata{Name: "no-tools"},
		Layers:   []capstack.Layer{{ID: "l1", Name: "Layer One", Order: 1}},
		Capabilities: []capstack.Capability{
			{ID: "a", Name: "A", LayerID: "l1"},
			{ID: "b", Name: "B", LayerID: "l1"},
		},
	}
	opts := DefaultSVGOptions()
	opts.Layout = SVGLayoutHub
	out, err := RenderSVGString(doc, opts)
	if err != nil {
		t.Fatalf("RenderSVGString: %v", err)
	}
	if !strings.Contains(out, "2 capabilities") {
		t.Errorf("expected capability-count fallback in node subtitle")
	}
}

func TestRenderSVGUnknownLayout(t *testing.T) {
	_, err := RenderSVGString(svgTestStack(), SVGOptions{Layout: "bogus", Theme: DefaultSVGTheme()})
	if err == nil {
		t.Fatal("expected error for unknown layout")
	}
}
