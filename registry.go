package errorlib

import "fmt"

// Registry holds a map of error code → AppError.
// Each service creates its own Registry and registers its own codes.
// The library provides this type so every service doesn't rewrite the same logic.
type Registry struct {
	errors map[string]AppError
}

// NewRegistry creates a new empty Registry.
func NewRegistry() *Registry {
	return &Registry{
		errors: make(map[string]AppError),
	}
}

// Register adds an AppError to the registry.
// Call this at service startup for each error code.
func (r *Registry) Register(err AppError) {
	r.errors[err.Code] = err
}

// Get retrieves an AppError by its code string.
// Returns an error if the code is not found.
func (r *Registry) Get(code string) (AppError, error) {
	err, found := r.errors[code]
	if !found {
		return AppError{}, fmt.Errorf("error code %q not found in registry", code)
	}
	return err, nil
}

// MustGet retrieves an AppError by its code string.
// Panics if the code is not found — use only when you are certain the code exists.
func (r *Registry) MustGet(code string) AppError {
	err, found := r.errors[code]
	if !found {
		panic("error code not found in registry: " + code)
	}
	return err
}
