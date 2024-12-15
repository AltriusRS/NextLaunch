package errors

import "fmt"

type ErrorType uint16

const (
	// Generic Errors

	// TypeUnknown is returned when an unknown error occurs.
	// This could be a serious error, but its cause is unknown
	// Severity: unknown
	TypeUnknown ErrorType = iota

	/*
		Config Errors
	*/
	// ConfigDirectoryNotFound is returned when the config directory
	// is not found
	// Severity: medium
	ConfigDirectoryNotFound

	/*
		Telemetry Errors
	*/
	// TelemetryDisabled is returned when the telemetry is disabled
	// in the config
	// Severity: low
	TelemetryDisabled

	// TelemetryLevelTooLow is returned when the telemetry level is
	// too low for the triggered event
	// Severity: low
	TelemetryLevelTooLow

	// TelemetryFeatureFlagNotFound is returned when the feature flag
	// is not found
	// Severity: low
	TelemetryFeatureFlagNotFound

	// TelemetryFeatureFlagEvaluationFailed is returned when the feature
	// flag evaluation failed
	// Severity: low
	TelemetryFeatureFlagEvaluationFailed

	// TelemetryTokenNotFound is returned when the telemetry client is called
	// with an empty token
	// Severity: low
	TelemetryTokenNotFound

	// TelemetryKeyNotFound is returned when the telemetry client is called
	// with an empty key
	// Severity: low
	TelemetryKeyNotFound

	// TelemetryDistinctIdentifierNotFound is returned when the telemetry client
	// cannot obtain a distinct identifier
	// Severity: low
	TelemetryDistinctIdentifierNotFound

	// TelemetryPosthogInitFailed is returned when the telemetry client fails
	// to initialize the posthog client
	// Severity: low
	TelemetryPosthogInitFailed

	// TelemetryPrometheusInitFailed is returned when the telemetry client fails
	// to initialize the prometheus exporter
	TelemetryPrometheusInitFailed

	//	TelemetryUniqueIdentifierNotFound is returned when the telemetry client
	// cannot obtain a unique identifier
	// Severity: low
	TelemetryUniqueIdentifierNotFound
)

// Array decoding the error type names to make it easier to find the error
// type by name
var errorTypeNames = map[ErrorType]string{
	// Generic Errors
	TypeUnknown: "Unknown",

	//	Config Errors
	ConfigDirectoryNotFound: "ConfigDirectoryNotFound",

	//	Telemetry Errors
	TelemetryDisabled:                    "TelemetryDisabled",
	TelemetryLevelTooLow:                 "TelemetryLevelTooLow",
	TelemetryFeatureFlagNotFound:         "TelemetryFeatureFlagNotFound",
	TelemetryFeatureFlagEvaluationFailed: "TelemetryFeatureFlagEvaluationFailed",
	TelemetryTokenNotFound:               "TelemetryTokenNotFound",
	TelemetryKeyNotFound:                 "TelemetryKeyNotFound",
	TelemetryDistinctIdentifierNotFound:  "TelemetryDistinctIdentifierNotFound",
	TelemetryPosthogInitFailed:           "TelemetryPosthogInitFailed",
	TelemetryPrometheusInitFailed:        "TelemetryPrometheusInitFailed",
	TelemetryUniqueIdentifierNotFound:    "TelemetryUniqueIdentifierNotFound",
}

type NextLaunchError struct {
	Code     ErrorType
	Message  string
	CodeName string
	Fatal    bool
}

func (e *NextLaunchError) Error() string {
	return fmt.Sprintf("%d - %s: %s", e.Code, e.CodeName, e.Message)
}

func (e *NextLaunchError) FatalError() {
	if e.Fatal {
		panic(e)
	}
}

func NewError(code ErrorType, error error, fatal bool) *NextLaunchError {
	return &NextLaunchError{
		Code:     code,
		Message:  error.Error(),
		CodeName: errorTypeNames[code],
		Fatal:    fatal,
	}
}

func NewErrorf(code ErrorType, fatal bool, format string, a ...interface{}) *NextLaunchError {
	return &NextLaunchError{
		Code:     code,
		Message:  fmt.Sprintf(format, a...),
		CodeName: errorTypeNames[code],
		Fatal:    fatal,
	}
}
