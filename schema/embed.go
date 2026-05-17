package schema

import (
	_ "embed"
	"encoding/json"
)

//go:embed capability-stack.schema.json
var capabilityStackSchema []byte

// CapabilityStackSchema returns the embedded JSON Schema for CapabilityStack documents.
func CapabilityStackSchema() []byte {
	return capabilityStackSchema
}

// CapabilityStackSchemaString returns the schema as a string.
func CapabilityStackSchemaString() string {
	return string(capabilityStackSchema)
}

// CapabilityStackSchemaMap returns the schema as a map.
func CapabilityStackSchemaMap() (map[string]any, error) {
	var m map[string]any
	if err := json.Unmarshal(capabilityStackSchema, &m); err != nil {
		return nil, err
	}
	return m, nil
}
