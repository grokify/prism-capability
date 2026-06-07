package capstack

import pf "github.com/grokify/priority-frameworks"

// Importance constants define static weights for categories, layers, and capabilities.
// These represent the inherent importance of "-ilities" (security, availability, etc.)
// and are used in conjunction with current state to calculate dynamic priority.
//
// Deprecated: Use priority-frameworks package directly for new code.
const (
	ImportanceCritical = "critical"
	ImportanceHigh     = "high"
	ImportanceMedium   = "medium"
	ImportanceLow      = "low"
)

// AllImportanceLevels returns all valid importance levels in descending order.
func AllImportanceLevels() []string {
	return []string{
		ImportanceCritical,
		ImportanceHigh,
		ImportanceMedium,
		ImportanceLow,
	}
}

// ImportanceWeight returns a numeric weight for the importance level.
// Higher weights indicate higher importance.
func ImportanceWeight(importance string) int {
	f := pf.Severity()
	idx := f.IndexOf(importance)
	if idx < 0 {
		return 2 // Default to medium
	}
	// Convert index to weight (0=highest -> 4, 4=lowest -> 0)
	// But we want critical=4, high=3, medium=2, low=1
	return len(f.Levels) - idx
}

// Priority constants define dynamic priority levels based on current state.
// These are calculated by combining importance with maturity gap.
const (
	PriorityP0 = "P0" // Immediate action required
	PriorityP1 = "P1" // High priority
	PriorityP2 = "P2" // Medium priority
	PriorityP3 = "P3" // Low priority
)

// AllPriorityLevels returns all valid priority levels in descending order.
func AllPriorityLevels() []string {
	return []string{
		PriorityP0,
		PriorityP1,
		PriorityP2,
		PriorityP3,
	}
}

// DynamicPriorityWeight returns a numeric weight for the dynamic priority level (P0-P3).
// Higher weights indicate higher priority.
func DynamicPriorityWeight(priority string) int {
	f := pf.Priority()
	idx := f.IndexOf(priority)
	if idx < 0 {
		return 2 // Default to P2
	}
	return len(f.Levels) - idx
}

// CalculatePriority determines dynamic priority based on importance and maturity gap.
// importance: the static importance level (critical, high, medium, low)
// currentLevel: current maturity level (1-5)
// targetLevel: target maturity level (1-5)
// Returns P0-P3 based on the combination.
func CalculatePriority(importance string, currentLevel, targetLevel int) string {
	if currentLevel >= targetLevel {
		return PriorityP3 // Already at or above target
	}

	gap := targetLevel - currentLevel
	weight := ImportanceWeight(importance)

	// Priority score: importance weight * gap
	score := weight * gap

	switch {
	case score >= 8: // Critical with 2+ gap, or High with 3+ gap
		return PriorityP0
	case score >= 4: // High with 2 gap, or Medium with 2+ gap
		return PriorityP1
	case score >= 2: // Any importance with small gap
		return PriorityP2
	default:
		return PriorityP3
	}
}

// SeverityFramework returns the Severity priority framework.
func SeverityFramework() *pf.Framework {
	return pf.Severity()
}

// PriorityFrameworkP returns the P# priority framework.
func PriorityFrameworkP() *pf.Framework {
	return pf.Priority()
}

// MoSCoWFramework returns the MoSCoW priority framework.
func MoSCoWFramework() *pf.Framework {
	return pf.MoSCoW()
}

// IETFFramework returns the IETF RFC 2119 requirement framework.
func IETFFramework() *pf.Framework {
	return pf.IETF()
}

// GeneralFramework returns the general requirement framework.
func GeneralFramework() *pf.Framework {
	return pf.General()
}

// GetFramework returns a priority framework by ID.
func GetFramework(id string) *pf.Framework {
	return pf.Get(id)
}
