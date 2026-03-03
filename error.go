package errorlib

// AppError is the standard error structure used across all microservices.
// Every service must use this shape — no custom error structs allowed.
type AppError struct {
	Code        string
	Description string
}

// Error implements the built-in Go error interface.
// This allows AppError to be used anywhere a standard error is expected.
func (e AppError) Error() string {
	return e.Code + ": " + e.Description
}
