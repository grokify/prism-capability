package capstack

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// ValidationError represents a validation error with context.
type ValidationError struct {
	Field   string
	Value   string
	Message string
}

func (e *ValidationError) Error() string {
	if e.Value != "" {
		return fmt.Sprintf("%s: %s (value: %q)", e.Field, e.Message, e.Value)
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors is a collection of validation errors.
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return ""
	}
	var msgs []string
	for _, e := range ve {
		msgs = append(msgs, e.Error())
	}
	return strings.Join(msgs, "; ")
}

// HasErrors returns true if there are any validation errors.
func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}

// kebabCasePattern matches valid kebab-case identifiers.
var kebabCasePattern = regexp.MustCompile(`^[a-z][a-z0-9]*(-[a-z0-9]+)*$`)

// ValidateKebabCase validates that a string is in kebab-case format.
func ValidateKebabCase(s string) bool {
	return kebabCasePattern.MatchString(s)
}

// ValidateCapabilityStatus validates a capability status value.
func ValidateCapabilityStatus(status string) error {
	if status == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllCapabilityStatuses(), status) {
		return fmt.Errorf("invalid status %q, must be one of: %s", status, strings.Join(AllCapabilityStatuses(), ", "))
	}
	return nil
}

// ValidatePriority validates a priority value.
func ValidatePriority(priority string) error {
	if priority == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllPriorities(), priority) {
		return fmt.Errorf("invalid priority %q, must be one of: %s", priority, strings.Join(AllPriorities(), ", "))
	}
	return nil
}

// ValidateDomain validates a domain value.
func ValidateDomain(domain string) error {
	if domain == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllDomains(), domain) {
		return fmt.Errorf("invalid domain %q, must be one of: %s", domain, strings.Join(AllDomains(), ", "))
	}
	return nil
}

// ValidatePhase validates a phase value.
func ValidatePhase(phase string) error {
	if phase == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllPhases(), phase) {
		return fmt.Errorf("invalid phase %q, must be one of: %s", phase, strings.Join(AllPhases(), ", "))
	}
	return nil
}

// ValidateNistCsfFunction validates a NIST CSF function value.
func ValidateNistCsfFunction(fn string) error {
	if fn == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllNistCsfFunctions(), fn) {
		return fmt.Errorf("invalid NIST CSF function %q, must be one of: %s", fn, strings.Join(AllNistCsfFunctions(), ", "))
	}
	return nil
}

// ValidateToolType validates a tool type value.
func ValidateToolType(toolType string) error {
	if toolType == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllToolTypes(), toolType) {
		return fmt.Errorf("invalid tool type %q, must be one of: %s", toolType, strings.Join(AllToolTypes(), ", "))
	}
	return nil
}

// ValidateToolStatus validates a tool status value.
func ValidateToolStatus(status string) error {
	if status == "" {
		return nil // Optional field
	}
	if !slices.Contains(AllToolStatuses(), status) {
		return fmt.Errorf("invalid tool status %q, must be one of: %s", status, strings.Join(AllToolStatuses(), ", "))
	}
	return nil
}

// ValidateFramework validates a framework value.
func ValidateFramework(framework string) error {
	if framework == "" {
		return fmt.Errorf("framework is required")
	}
	if !slices.Contains(AllFrameworks(), framework) {
		return fmt.Errorf("invalid framework %q, must be one of: %s", framework, strings.Join(AllFrameworks(), ", "))
	}
	return nil
}

// Validate validates a Tool and returns validation errors.
func (t *Tool) Validate() ValidationErrors {
	var errs ValidationErrors

	if t.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	if err := ValidateToolType(t.Type); err != nil {
		errs = append(errs, ValidationError{Field: "type", Value: t.Type, Message: err.Error()})
	}

	if err := ValidateToolStatus(t.Status); err != nil {
		errs = append(errs, ValidationError{Field: "status", Value: t.Status, Message: err.Error()})
	}

	return errs
}

// Validate validates a FrameworkMapping and returns validation errors.
func (fm *FrameworkMapping) Validate() ValidationErrors {
	var errs ValidationErrors

	if err := ValidateFramework(fm.Framework); err != nil {
		errs = append(errs, ValidationError{Field: "framework", Value: fm.Framework, Message: err.Error()})
	}

	if len(fm.Controls) == 0 {
		errs = append(errs, ValidationError{Field: "controls", Message: "at least one control is required"})
	}

	return errs
}

// Validate validates a Layer and returns validation errors.
func (l *Layer) Validate() ValidationErrors {
	var errs ValidationErrors

	if l.ID == "" {
		errs = append(errs, ValidationError{Field: "id", Message: "is required"})
	} else if !ValidateKebabCase(l.ID) {
		errs = append(errs, ValidationError{Field: "id", Value: l.ID, Message: "must be kebab-case"})
	}

	if l.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	if err := ValidatePhase(l.Phase); err != nil {
		errs = append(errs, ValidationError{Field: "phase", Value: l.Phase, Message: err.Error()})
	}

	if err := ValidateNistCsfFunction(l.NistCsfFunction); err != nil {
		errs = append(errs, ValidationError{Field: "nistCsfFunction", Value: l.NistCsfFunction, Message: err.Error()})
	}

	return errs
}

// Validate validates a Category and returns validation errors.
func (c *Category) Validate() ValidationErrors {
	var errs ValidationErrors

	if c.ID == "" {
		errs = append(errs, ValidationError{Field: "id", Message: "is required"})
	} else if !ValidateKebabCase(c.ID) {
		errs = append(errs, ValidationError{Field: "id", Value: c.ID, Message: "must be kebab-case"})
	}

	if c.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	return errs
}

// Validate validates a Capability and returns validation errors.
func (cap *Capability) Validate() ValidationErrors {
	var errs ValidationErrors

	if cap.ID == "" {
		errs = append(errs, ValidationError{Field: "id", Message: "is required"})
	} else if !ValidateKebabCase(cap.ID) {
		errs = append(errs, ValidationError{Field: "id", Value: cap.ID, Message: "must be kebab-case"})
	}

	if cap.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	if cap.LayerID == "" {
		errs = append(errs, ValidationError{Field: "layerId", Message: "is required"})
	}

	if err := ValidateCapabilityStatus(cap.Status); err != nil {
		errs = append(errs, ValidationError{Field: "status", Value: cap.Status, Message: err.Error()})
	}

	if err := ValidatePriority(cap.Priority); err != nil {
		errs = append(errs, ValidationError{Field: "priority", Value: cap.Priority, Message: err.Error()})
	}

	// Validate tags are kebab-case
	for i, tag := range cap.Tags {
		if !ValidateKebabCase(tag) {
			errs = append(errs, ValidationError{
				Field:   fmt.Sprintf("tags[%d]", i),
				Value:   tag,
				Message: "must be kebab-case",
			})
		}
	}

	// Validate tooling
	for i, tool := range cap.Tooling {
		toolErrs := tool.Validate()
		for _, e := range toolErrs {
			e.Field = fmt.Sprintf("tooling[%d].%s", i, e.Field)
			errs = append(errs, e)
		}
	}

	// Validate framework mappings
	for i, fm := range cap.FrameworkMappings {
		fmErrs := fm.Validate()
		for _, e := range fmErrs {
			e.Field = fmt.Sprintf("frameworkMappings[%d].%s", i, e.Field)
			errs = append(errs, e)
		}
	}

	return errs
}

// Validate validates a Metadata and returns validation errors.
func (m *Metadata) Validate() ValidationErrors {
	var errs ValidationErrors

	if m.Name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "is required"})
	}

	if m.Version == "" {
		errs = append(errs, ValidationError{Field: "version", Message: "is required"})
	}

	if err := ValidateDomain(m.Domain); err != nil {
		errs = append(errs, ValidationError{Field: "domain", Value: m.Domain, Message: err.Error()})
	}

	return errs
}

// Validate validates the entire CapabilityStack document.
func (cs *CapabilityStack) Validate() ValidationErrors {
	var errs ValidationErrors

	// Validate metadata
	metaErrs := cs.Metadata.Validate()
	for _, e := range metaErrs {
		e.Field = "metadata." + e.Field
		errs = append(errs, e)
	}

	// Validate layers
	if len(cs.Layers) == 0 {
		errs = append(errs, ValidationError{Field: "layers", Message: "at least one layer is required"})
	}

	seenLayerIDs := make(map[string]int)
	for i, layer := range cs.Layers {
		layerErrs := layer.Validate()
		for _, e := range layerErrs {
			e.Field = fmt.Sprintf("layers[%d].%s", i, e.Field)
			errs = append(errs, e)
		}

		// Check for duplicate layer IDs
		if layer.ID != "" {
			if prevIdx, exists := seenLayerIDs[layer.ID]; exists {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("layers[%d].id", i),
					Value:   layer.ID,
					Message: fmt.Sprintf("duplicate ID, also used at layers[%d]", prevIdx),
				})
			}
			seenLayerIDs[layer.ID] = i
		}
	}

	// Validate categories
	seenCategoryIDs := make(map[string]int)
	for i, cat := range cs.Categories {
		catErrs := cat.Validate()
		for _, e := range catErrs {
			e.Field = fmt.Sprintf("categories[%d].%s", i, e.Field)
			errs = append(errs, e)
		}

		// Check for duplicate category IDs
		if cat.ID != "" {
			if prevIdx, exists := seenCategoryIDs[cat.ID]; exists {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("categories[%d].id", i),
					Value:   cat.ID,
					Message: fmt.Sprintf("duplicate ID, also used at categories[%d]", prevIdx),
				})
			}
			seenCategoryIDs[cat.ID] = i
		}
	}

	// Validate capabilities
	if len(cs.Capabilities) == 0 {
		errs = append(errs, ValidationError{Field: "capabilities", Message: "at least one capability is required"})
	}

	seenCapIDs := make(map[string]int)
	allCaps := cs.AllCapabilities()
	for i, cap := range allCaps {
		var prefix string
		if i < len(cs.Capabilities) {
			prefix = fmt.Sprintf("capabilities[%d]", i)
		} else {
			prefix = fmt.Sprintf("foundational[%d]", i-len(cs.Capabilities))
		}

		capErrs := cap.Validate()
		for _, e := range capErrs {
			e.Field = prefix + "." + e.Field
			errs = append(errs, e)
		}

		// Check for duplicate capability IDs
		if cap.ID != "" {
			if prevIdx, exists := seenCapIDs[cap.ID]; exists {
				var prevPrefix string
				if prevIdx < len(cs.Capabilities) {
					prevPrefix = fmt.Sprintf("capabilities[%d]", prevIdx)
				} else {
					prevPrefix = fmt.Sprintf("foundational[%d]", prevIdx-len(cs.Capabilities))
				}
				errs = append(errs, ValidationError{
					Field:   prefix + ".id",
					Value:   cap.ID,
					Message: fmt.Sprintf("duplicate ID, also used at %s", prevPrefix),
				})
			}
			seenCapIDs[cap.ID] = i
		}

		// Validate layerId reference
		if cap.LayerID != "" && cs.GetLayerByID(cap.LayerID) == nil {
			errs = append(errs, ValidationError{
				Field:   prefix + ".layerId",
				Value:   cap.LayerID,
				Message: "references non-existent layer ID",
			})
		}

		// Validate categoryId reference
		if cap.CategoryID != "" && cs.GetCategoryByID(cap.CategoryID) == nil {
			errs = append(errs, ValidationError{
				Field:   prefix + ".categoryId",
				Value:   cap.CategoryID,
				Message: "references non-existent category ID",
			})
		}

		// Validate dependency references
		for j, depID := range cap.Dependencies {
			if cs.GetCapabilityByID(depID) == nil {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("%s.dependencies[%d]", prefix, j),
					Value:   depID,
					Message: "references non-existent capability ID",
				})
			}
		}

		// Validate enables references
		for j, enablesID := range cap.Enables {
			if cs.GetCapabilityByID(enablesID) == nil {
				errs = append(errs, ValidationError{
					Field:   fmt.Sprintf("%s.enables[%d]", prefix, j),
					Value:   enablesID,
					Message: "references non-existent capability ID",
				})
			}
		}
	}

	// Detect dependency cycles
	cycleErrs := cs.detectDependencyCycles()
	errs = append(errs, cycleErrs...)

	// Validate capability Order uniqueness (only if any non-zero Order exists)
	orderErrs := cs.validateCapabilityOrder()
	errs = append(errs, orderErrs...)

	return errs
}

// validateCapabilityOrder checks that Order values are globally unique when any non-zero Order exists.
func (cs *CapabilityStack) validateCapabilityOrder() ValidationErrors {
	var errs ValidationErrors

	// Check if any capability has explicit ordering
	hasExplicitOrder := false
	for _, cap := range cs.AllCapabilities() {
		if cap.Order != 0 {
			hasExplicitOrder = true
			break
		}
	}

	if !hasExplicitOrder {
		return nil // No explicit ordering, skip validation
	}

	// Check for duplicate Order values
	type orderInfo struct {
		capID  string
		prefix string
	}
	seen := make(map[int]orderInfo)
	allCaps := cs.AllCapabilities()

	for i, cap := range allCaps {
		var prefix string
		if i < len(cs.Capabilities) {
			prefix = fmt.Sprintf("capabilities[%d]", i)
		} else {
			prefix = fmt.Sprintf("foundational[%d]", i-len(cs.Capabilities))
		}

		if existing, ok := seen[cap.Order]; ok {
			errs = append(errs, ValidationError{
				Field:   prefix + ".order",
				Value:   fmt.Sprintf("%d", cap.Order),
				Message: fmt.Sprintf("duplicate order value, also used by %s (%s)", existing.capID, existing.prefix),
			})
		}
		seen[cap.Order] = orderInfo{capID: cap.ID, prefix: prefix}
	}

	return errs
}

// detectDependencyCycles checks for circular dependencies.
func (cs *CapabilityStack) detectDependencyCycles() ValidationErrors {
	var errs ValidationErrors

	// Build adjacency list
	deps := make(map[string][]string)
	for _, cap := range cs.AllCapabilities() {
		deps[cap.ID] = cap.Dependencies
	}

	// DFS-based cycle detection
	white := make(map[string]bool) // unvisited
	gray := make(map[string]bool)  // in current path
	black := make(map[string]bool) // fully processed

	for _, cap := range cs.AllCapabilities() {
		white[cap.ID] = true
	}

	var dfs func(id string, path []string) bool
	dfs = func(id string, path []string) bool {
		if black[id] {
			return false
		}
		if gray[id] {
			// Found cycle
			cycleStart := 0
			for i, p := range path {
				if p == id {
					cycleStart = i
					break
				}
			}
			cycle := append(path[cycleStart:], id)
			errs = append(errs, ValidationError{
				Field:   "dependencies",
				Value:   strings.Join(cycle, " -> "),
				Message: "circular dependency detected",
			})
			return true
		}

		delete(white, id)
		gray[id] = true

		for _, depID := range deps[id] {
			if dfs(depID, append(path, id)) {
				return true
			}
		}

		delete(gray, id)
		black[id] = true
		return false
	}

	for id := range white {
		if dfs(id, nil) {
			break // Stop after finding first cycle
		}
	}

	return errs
}
