package stores

type Cache interface {
	Put() error
	Get() (string, error)
	Has() bool
	Delete() (bool, error)
	Pull() (string, error)
}

// Put a value
func Put(c Cache) error {
	return c.Put()
}

// Retrieve a cached Value
func Get(c Cache) (string, error) {
	return c.Get()
}

// Check is key exists
func Has(c Cache) bool {
	return c.Has()
}

// Retrieve and delete value
func Pull(c Cache) (bool, error) {
	return c.Pull()
}
