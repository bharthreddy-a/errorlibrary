package errorlib

// LoadByArray bulk-registers a slice of AppErrors into a Registry.
// Each service calls this at startup to register all its error codes at once.
//
// Example usage in a service:
//
//	var Registry = errorlib.NewRegistry()
//
//	func init() {
//	    errorlib.Load(Registry, AllErrors)
//	}
func (r *Registry) LoadByArray(errors []AppError) {
	if r == nil || errors == nil {
		return
	}
	for _, err := range errors {
		r.Register(err)
	}
}
