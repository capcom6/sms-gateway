package smsgateway

type HealthStatus string

const (
	HealthStatusPass HealthStatus = "pass"
	HealthStatusWarn HealthStatus = "warn"
	HealthStatusFail HealthStatus = "fail"
)

// Details of a health check.
type HealthCheck struct {
	// A human-readable description of the check.
	Description string `json:"description,omitempty"`
	// Unit of measurement for the observed value.
	ObservedUnit string `json:"observedUnit,omitempty"`
	// Observed value of the check.
	ObservedValue int `json:"observedValue"`
	// Status of the check.
	// It can be one of the following values: "pass", "warn", or "fail".
	Status HealthStatus `json:"status"`
}

// Map of check names to their respective details.
type HealthChecks map[string]HealthCheck

// Health status of the application.
type HealthResponse struct {
	// Overall status of the application.
	// It can be one of the following values: "pass", "warn", or "fail".
	Status HealthStatus `json:"status"`
	// Version of the application.
	Version string `json:"version,omitempty"`
	// Release ID of the application.
	// It is used to identify the version of the application.
	ReleaseID int `json:"releaseId,omitempty"`
	// A map of check names to their respective details.
	Checks HealthChecks `json:"checks,omitempty"`
}
