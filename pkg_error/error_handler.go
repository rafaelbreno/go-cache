package pkg_error

import (
	"fmt"
	"runtime"
)

// Go Cache main error handler
type PkgError struct {
	Error   error
	Message string
	Tracer  Tracer
	Nil     bool
}

// Tracer struct shows important info for debugging
// E.g File's name, Line number, Function's name
type Tracer struct {
	File     string
	Line     int
	Function string
}

// Returns a nil Error
func NewNilError() PkgError {
	return PkgError{
		Nil: true,
	}
}

// Returns new instance of PkgError
func NewError(err error) PkgError {
	pkgErr := PkgError{
		Error: err,
		Nil:   false,
	}.setTracer()

	return pkgErr
}

func (p PkgError) SetMessage(msgError string, args ...interface{}) PkgError {
	msgFmt := fmt.Sprintf(msgError, args...)

	p.Message = msgFmt

	return p
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
