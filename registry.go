package errorlib

import (
	"fmt"
	"sync"
)

// Registry holds a map of error code → AppError.
// Each service creates its own Registry and registers its own codes.
// The library provides this type so every service doesn't rewrite the same logic.
type Registry struct {
	mu     sync.RWMutex
	errors map[string]AppError
}

var (
	instance *Registry
	once     sync.Once
)

func GetRegistry() *Registry {
	once.Do(func() {
		instance = &Registry{
			errors: make(map[string]AppError),
		}
	})
	return instance
}

// Register adds an AppError to the registry.
// Call this at service startup for each error code.
func (r *Registry) Register(err AppError) {
	if err.Code == "" || err.Description == "" {
		//code and description cannot equal to empty string
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.errors[err.Code] = err
}

// Get retrieves an AppError by its code string.
// Returns an error if the code is not found.
func (r *Registry) Get(code string) (AppError, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	err, found := r.errors[code]
	if !found {
		return AppError{}, fmt.Errorf("error code %q not found in registry", code)
	}
	return err, nil
}
