package capstack

import (
	"testing"
)

func TestAllCapabilityStatuses(t *testing.T) {
	statuses := AllCapabilityStatuses()
	if len(statuses) != 5 {
		t.Errorf("Expected 5 statuses, got %d", len(statuses))
	}

	expected := []string{"planned", "in-progress", "implemented", "operational", "deprecated"}
	for i, s := range expected {
		if statuses[i] != s {
			t.Errorf("Status[%d] = %q, want %q", i, statuses[i], s)
		}
	}
}

func TestAllPriorities(t *testing.T) {
	priorities := AllPriorities()
	if len(priorities) != 4 {
		t.Errorf("Expected 4 priorities, got %d", len(priorities))
	}

	expected := []string{"critical", "high", "medium", "low"}
	for i, p := range expected {
		if priorities[i] != p {
			t.Errorf("Priority[%d] = %q, want %q", i, priorities[i], p)
		}
	}
}

func TestAllDomains(t *testing.T) {
	domains := AllDomains()
	if len(domains) != 8 {
		t.Errorf("Expected 8 domains, got %d", len(domains))
	}

	// Verify key domains exist
	found := make(map[string]bool)
	for _, d := range domains {
		found[d] = true
	}

	required := []string{"security", "operations", "ai", "data"}
	for _, r := range required {
		if !found[r] {
			t.Errorf("Missing required domain: %s", r)
		}
	}
}

func TestAllPhases(t *testing.T) {
	phases := AllPhases()
	if len(phases) != 10 {
		t.Errorf("Expected 10 phases, got %d", len(phases))
	}
}

func TestAllNistCsfFunctions(t *testing.T) {
	functions := AllNistCsfFunctions()
	if len(functions) != 6 {
		t.Errorf("Expected 6 NIST CSF functions, got %d", len(functions))
	}

	expected := []string{"govern", "identify", "protect", "detect", "respond", "recover"}
	for i, f := range expected {
		if functions[i] != f {
			t.Errorf("Function[%d] = %q, want %q", i, functions[i], f)
		}
	}
}

func TestAllToolTypes(t *testing.T) {
	types := AllToolTypes()
	if len(types) != 4 {
		t.Errorf("Expected 4 tool types, got %d", len(types))
	}
}

func TestAllToolStatuses(t *testing.T) {
	statuses := AllToolStatuses()
	if len(statuses) != 4 {
		t.Errorf("Expected 4 tool statuses, got %d", len(statuses))
	}
}

func TestAllFrameworks(t *testing.T) {
	frameworks := AllFrameworks()
	if len(frameworks) != 10 {
		t.Errorf("Expected 10 frameworks, got %d", len(frameworks))
	}

	// Verify key frameworks exist
	found := make(map[string]bool)
	for _, f := range frameworks {
		found[f] = true
	}

	required := []string{"nist-csf-2.0", "iso-27001", "soc2", "slsa"}
	for _, r := range required {
		if !found[r] {
			t.Errorf("Missing required framework: %s", r)
		}
	}
}
