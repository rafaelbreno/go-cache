package stores

import (
	"github.com/rafaelbreno/go-cache/pkg_error"
)

type CacheInterface interface {
	Put() pkg_error.PkgError
	Get() (string, pkg_error.PkgError)
	Has() (bool, pkg_error.PkgError)
	Delete() pkg_error.PkgError
	Pull() (string, pkg_error.PkgError)
}

// Put a value
func Put(c CacheInterface) pkg_error.PkgError {
	return c.Put()
}

// Retrieve a cached Value
func Get(c CacheInterface) (string, pkg_error.PkgError) {
	return c.Get()
}

// Check is key exists
func Has(c CacheInterface) (bool, pkg_error.PkgError) {
	return c.Has()
}

// Check is key exists
func Delete(c CacheInterface) pkg_error.PkgError {
	return c.Delete()
}

// Retrieve and delete value
func Pull(c CacheInterface) (string, pkg_error.PkgError) {
	return c.Pull()
}
