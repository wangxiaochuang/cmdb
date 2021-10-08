package errors

var defaultHandler ErrorsInterface = &pkgError{}

// SetDefaultHandler set error handler
func SetDefaultHandler(handler ErrorsInterface) {
	defaultHandler = handler
}

// New returns new error with message
var New = defaultHandler.New()
