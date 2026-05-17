package capstack

// CapabilityStatus constants represent the lifecycle status of a capability.
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

// Priority constants represent implementation priority levels.
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

// Domain constants represent primary capability stack domains.
const (
	DomainSecurity       = "security"
	DomainAI             = "ai"
	DomainPlatform       = "platform"
	DomainData           = "data"
	DomainObservability  = "observability"
	DomainInfrastructure = "infrastructure"
	DomainProduct        = "product"
	DomainOperations     = "operations"
)

// AllDomains returns all valid domain values.
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

// Phase constants represent SDLC or lifecycle phases.
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

// NistCsfFunction constants represent NIST CSF 2.0 functions.
const (
	NistCsfGovern   = "govern"
	NistCsfIdentify = "identify"
	NistCsfProtect  = "protect"
	NistCsfDetect   = "detect"
	NistCsfRespond  = "respond"
	NistCsfRecover  = "recover"
)

// AllNistCsfFunctions returns all valid NIST CSF function values.
func AllNistCsfFunctions() []string {
	return []string{
		NistCsfGovern,
		NistCsfIdentify,
		NistCsfProtect,
		NistCsfDetect,
		NistCsfRespond,
		NistCsfRecover,
	}
}

// ToolType constants represent tool/product types.
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
	}
}
