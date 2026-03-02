package errorlib

// AppError is the standard error structure used across all microservices.
// Every service must use this shape — no custom error structs allowed.
type AppError struct {
	Code        string // e.g. "AUTH_001", "PAY_002", "CMN_001"
	ServiceName string // e.g. "auth-service", "payment-service"
	Description string // e.g. "Invalid or expired token"
}

// Error implements the built-in Go error interface.
// This allows AppError to be used anywhere a standard error is expected.
func (e AppError) Error() string {
	return "[" + e.ServiceName + "] " + e.Code + ": " + e.Description
}
