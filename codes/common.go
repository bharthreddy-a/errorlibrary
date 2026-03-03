package codes

import "errorlib"

// Common errors shared across ALL services.
// Prefix: CMN
// These are owned by the sre team.
// To add a new common error, raise a PR to the errorlib repo.
var (
	ErrInternalServer = errorlib.AppError{
		Code:        "CMN_001",
		Description: "Internal server error",
	}

	ErrServiceUnavailable = errorlib.AppError{
		Code:        "CMN_002",
		Description: "Service temporarily unavailable",
	}

	ErrRequestTimeout = errorlib.AppError{
		Code:        "CMN_003",
		Description: "Request timed out",
	}

	ErrUnauthorized = errorlib.AppError{
		Code:        "CMN_004",
		Description: "Unauthorized access",
	}

	ErrNotFound = errorlib.AppError{
		Code:        "CMN_005",
		Description: "Resource not found",
	}
)

// AllCommonErrors is the full list of common errors.
// Services load this into their registry at startup alongside their own errors.
var AllCommonErrors = []errorlib.AppError{
	ErrInternalServer,
	ErrServiceUnavailable,
	ErrRequestTimeout,
	ErrUnauthorized,
	ErrNotFound,
}
