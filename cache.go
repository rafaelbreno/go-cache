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

	c.SetPath()

	return c, nil
}

func (c *Cache) SetPath() {
	keyBytes := sha1.Sum([]byte(c.key))

	c.Path = fmt.Sprintf("/%x/%x/", keyBytes[0], keyBytes[1])

	for i := 0; i < len(keyBytes); i++ {
		c.Path += fmt.Sprintf("%x", keyBytes[i])
	}
}
