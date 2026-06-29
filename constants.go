package capstack

import core "github.com/grokify/prism-core"

// CapabilityStatus constants represent the lifecycle status of a capability.
// Note: Uses hyphen format ("in-progress") for JSON/YAML compatibility.
// prism-core uses underscore format - these remain local for backward compatibility.
const (
	StatusPlanned     = "planned"
	StatusInProgress  = "in-progress"
	StatusImplemented = "implemented"
	StatusOperational = "operational"
	StatusDeprecated  = "deprecated"
)

// AllCapabilityStatuses returns all valid capability status values.
func AllCapabilityStatuses() []string {
	return []string{
		StatusPlanned,
		StatusInProgress,
		StatusImplemented,
		StatusOperational,
		StatusDeprecated,
	}
}

// Priority constants for severity-based prioritization.
// Now backed by priority-frameworks package.
const (
	PriorityCritical = "critical"
	PriorityHigh     = "high"
	PriorityMedium   = "medium"
	PriorityLow      = "low"
)

// AllPriorities returns all valid priority values.
func AllPriorities() []string {
	return []string{
		PriorityCritical,
		PriorityHigh,
		PriorityMedium,
		PriorityLow,
	}
}

// ValidPriority checks if a priority value is valid.
func ValidPriority(priority string) bool {
	f := SeverityFramework()
	return f.IndexOf(priority) >= 0 || priority == ""
}

// PriorityWeight returns a numeric weight for sorting priorities.
// Higher weight = higher priority.
func PriorityWeight(priority string) int {
	f := SeverityFramework()
	idx := f.IndexOf(priority)
	if idx < 0 {
		return 0
	}
	return len(f.Levels) - idx
}

// Domain constants imported from prism-core.
const (
	DomainSecurity       = core.DomainSecurity
	DomainAI             = core.DomainAI
	DomainPlatform       = core.DomainPlatform
	DomainData           = core.DomainData
	DomainObservability  = core.DomainObservability
	DomainInfrastructure = core.DomainInfrastructure
	DomainProduct        = core.DomainProduct
	DomainOperations     = core.DomainOperations
)

// AllDomains returns all valid domain values for capability stacks.
// Note: Returns the 8 domains used by this module, not all prism-core domains.
func AllDomains() []string {
	return []string{
		DomainSecurity,
		DomainAI,
		DomainPlatform,
		DomainData,
		DomainObservability,
		DomainInfrastructure,
		DomainProduct,
		DomainOperations,
	}
}

// ValidDomain checks if a domain value is valid.
func ValidDomain(domain string) bool {
	return core.ValidDomain(domain)
}

// DomainDisplayName returns a human-readable name for a domain.
func DomainDisplayName(domain string) string {
	return core.DomainDisplayName(domain)
}

// Phase constants represent SDLC or lifecycle phases.
// These are more granular than prism-core stages and remain local.
const (
	PhasePlan    = "plan"
	PhaseDesign  = "design"
	PhaseBuild   = "build"
	PhaseTest    = "test"
	PhaseRelease = "release"
	PhaseDeploy  = "deploy"
	PhaseOperate = "operate"
	PhaseMonitor = "monitor"
	PhaseRespond = "respond"
	PhaseRecover = "recover"
)

// AllPhases returns all valid phase values.
func AllPhases() []string {
	return []string{
		PhasePlan,
		PhaseDesign,
		PhaseBuild,
		PhaseTest,
		PhaseRelease,
		PhaseDeploy,
		PhaseOperate,
		PhaseMonitor,
		PhaseRespond,
		PhaseRecover,
	}
}

// NIST CSF function constants imported from prism-core.
// Local names use NistCsf prefix for backward compatibility.
const (
	NistCsfGovern   = core.NISTCSFGovern
	NistCsfIdentify = core.NISTCSFIdentify
	NistCsfProtect  = core.NISTCSFProtect
	NistCsfDetect   = core.NISTCSFDetect
	NistCsfRespond  = core.NISTCSFRespond
	NistCsfRecover  = core.NISTCSFRecover
)

// AllNistCsfFunctions returns all valid NIST CSF function values.
func AllNistCsfFunctions() []string {
	return core.NISTCSFFunctions()
}

// NISTCSFFunctionSortWeight returns a sort weight for NIST CSF functions.
func NISTCSFFunctionSortWeight(function string) int {
	return core.NISTCSFFunctionSortWeight(function)
}

// ToolType constants represent tool/product types.
// These are capability-specific and remain local.
const (
	ToolTypeCommercial     = "commercial"
	ToolTypeOpenSource     = "open-source"
	ToolTypeInternal       = "internal"
	ToolTypeManagedService = "managed-service"
)

// AllToolTypes returns all valid tool type values.
func AllToolTypes() []string {
	return []string{
		ToolTypeCommercial,
		ToolTypeOpenSource,
		ToolTypeInternal,
		ToolTypeManagedService,
	}
}

// ToolStatus constants represent tool deployment status.
const (
	ToolStatusEvaluating = "evaluating"
	ToolStatusPiloting   = "piloting"
	ToolStatusDeployed   = "deployed"
	ToolStatusDeprecated = "deprecated"
)

// AllToolStatuses returns all valid tool status values.
func AllToolStatuses() []string {
	return []string{
		ToolStatusEvaluating,
		ToolStatusPiloting,
		ToolStatusDeployed,
		ToolStatusDeprecated,
	}
}

// Framework constants represent compliance/security frameworks.
// Note: Uses kebab-case format for JSON/YAML compatibility.
// prism-core uses UPPER_SNAKE format - these remain local for backward compatibility.
const (
	FrameworkNISTCSF2    = "nist-csf-2.0"
	FrameworkNIST80053   = "nist-800-53"
	FrameworkISO27001    = "iso-27001"
	FrameworkSOC2        = "soc2"
	FrameworkPCIDSS      = "pci-dss"
	FrameworkCIS         = "cis"
	FrameworkMITREATTACK = "mitre-attack"
	FrameworkOWASP       = "owasp"
	FrameworkSLSA        = "slsa"
	FrameworkSSDF        = "ssdf"

	// Operations and DevOps Frameworks
	FrameworkDORA = "dora" // DevOps Research and Assessment
	FrameworkSRE  = "sre"  // Site Reliability Engineering

	// Developer Productivity Frameworks
	FrameworkSPACE   = "space"    // Satisfaction, Performance, Activity, Communication, Efficiency
	FrameworkAIDORA  = "ai-dora"  // AI-Modified DevOps Research and Assessment
	FrameworkAISPACE = "ai-space" // AI-Modified Developer Productivity Framework
)

// AllFrameworks returns all valid framework values.
func AllFrameworks() []string {
	return []string{
		FrameworkNISTCSF2,
		FrameworkNIST80053,
		FrameworkISO27001,
		FrameworkSOC2,
		FrameworkPCIDSS,
		FrameworkCIS,
		FrameworkMITREATTACK,
		FrameworkOWASP,
		FrameworkSLSA,
		FrameworkSSDF,
		FrameworkDORA,
		FrameworkSRE,
		FrameworkSPACE,
		FrameworkAIDORA,
		FrameworkAISPACE,
	}
}

// OperationsFrameworks returns operations-focused frameworks.
func OperationsFrameworks() []string {
	return []string{
		FrameworkDORA,
		FrameworkSRE,
		FrameworkAIDORA,
	}
}

// DeveloperProductivityFrameworks returns developer productivity frameworks.
func DeveloperProductivityFrameworks() []string {
	return []string{
		FrameworkSPACE,
		FrameworkAISPACE,
	}
}

// AIFrameworks returns AI-enhanced frameworks.
func AIFrameworks() []string {
	return []string{
		FrameworkAIDORA,
		FrameworkAISPACE,
	}
}
