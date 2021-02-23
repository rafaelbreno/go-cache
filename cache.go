package cache

import (
	"crypto/sha1"
	"fmt"
)

type Cache struct {
	key   string
	value string
	Path  string
}

func Put(key, value string) (Cache, error) {
	if key == "" {
		return Cache{}, fmt.Errorf("'key' must not be nil")
	}

	if value == "" {
		return Cache{}, fmt.Errorf("'value' must not be nil")
	}

	c := Cache{
		key:   key,
		value: value,
	}

	if err := c.SetPath(); err != nil {
		return Cache{}, err
	}

	return c, nil
}

func (c *Cache) SetPath() error {
	keyBytes := sha1.Sum([]byte(c.key))

	if len(keyBytes) == 0 {
		return fmt.Errorf("Error converting key to sha1 bytes array")
	}

	c.Path = fmt.Sprintf("/%x/%x/", keyBytes[0], keyBytes[1])

	for i := 0; i < len(keyBytes); i++ {
		c.Path += fmt.Sprintf("%x", keyBytes[i])
	}

	return nil
}
