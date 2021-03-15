package pkg_error

import "runtime"

// Go Cache main error handler
type PkgError struct {
	Error   error
	Message string
	Tracer  Tracer
}

// Tracer struct shows important info for debugging
// E.g File's name, Line number, Function's name
type Tracer struct {
	File     string
	Line     int
	Function string
}

// Returns new instance of PkgError
func NewPkgError(err error, msg string) PkgError {
	pkgErr := PkgError{
		Error:   err,
		Message: msg,
	}.setTracer()

	return pkgErr
}

// Set tracer data
func (p PkgError) setTracer() PkgError {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	t := Tracer{
		File:     frame.File,
		Line:     frame.Line,
		Function: frame.Function,
	}

	p.Tracer = t

	return p
}

// Check if error is nil
func (p PkgError) IsNil() bool {
	return p.Error == nil
}
