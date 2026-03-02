package errorlib

// Load bulk-registers a slice of AppErrors into a Registry.
// Each service calls this at startup to register all its error codes at once.
//
// Example usage in a service:
//
//	var Registry = errorlib.NewRegistry()
//
//	func init() {
//	    errorlib.Load(Registry, AllErrors)
//	}
func Load(r *Registry, errors []AppError) {
	if r == nil {
		return
	}
	for _, err := range errors {
		r.Register(err)
	}
}
