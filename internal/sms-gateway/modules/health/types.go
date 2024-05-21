package health

import "context"

type Status string
type statusLevel int

const (
	StatusPass Status = "pass"
	StatusWarn Status = "warn"
	StatusFail Status = "fail"

	levelPass statusLevel = 0
	levelWarn statusLevel = 1
	levelFail statusLevel = 2
)

var statusLevels = map[statusLevel]Status{
	levelPass: StatusPass,
	levelWarn: StatusWarn,
	levelFail: StatusFail,
}

// Health status of the application.
type Check struct {
	// Overall status of the application.
	// It can be one of the following values: "pass", "warn", or "fail".
	Status Status
	// A map of check names to their respective details.
	Checks Checks
}

// Details of a health check.
type CheckDetail struct {
	// A human-readable description of the check.
	Description string
	// Unit of measurement for the observed value.
	ObservedUnit string
	// Observed value of the check.
	ObservedValue int
	// Status of the check.
	// It can be one of the following values: "pass", "warn", or "fail".
	Status Status
}

// Map of check names to their respective details.
type Checks map[string]CheckDetail

type HealthProvider interface {
	Name() string
	HealthCheck(ctx context.Context) (Checks, error)
}
