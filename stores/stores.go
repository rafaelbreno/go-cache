package stores

type CacheInterface interface {
	Put() error
	Get() (string, error)
	Has() bool
	Delete() error
	Pull() (string, error)
}

// Put a value
func Put(c CacheInterface) error {
	return c.Put()
}

// Retrieve a cached Value
func Get(c CacheInterface) (string, error) {
	return c.Get()
}

// Check is key exists
func Has(c CacheInterface) bool {
	return c.Has()
}

// Check is key exists
func Delete(c CacheInterface) error {
	return c.Delete()
}

// Retrieve and delete value
func Pull(c CacheInterface) (string, error) {
	return c.Pull()
}
