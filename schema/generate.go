//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"

	capstack "github.com/grokify/prism-capability"
	"github.com/invopop/jsonschema"
)

func main() {
	r := new(jsonschema.Reflector)
	r.ExpandedStruct = true

	schema := r.Reflect(&capstack.CapabilityStack{})
	schema.ID = "https://github.com/grokify/prism-capability/schema/capability-stack.schema.json"
	schema.Title = "Capability Stack Specification"
	schema.Description = "A specification for defining capability stacks with layers, capabilities, and PRISM maturity integration"

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling schema: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("schema/capability-stack.schema.json", data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Generated schema/capability-stack.schema.json")
}
