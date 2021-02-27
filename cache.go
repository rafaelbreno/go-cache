package cache

import (
	"crypto/sha1"
	"fmt"
	"os"
)

type Cache struct {
	key      string // Cache identifier - Cannot be null
	value    []byte // Cache value itself
	Path     string // Cache file path with the stored value
	FileName string // Cache file's name
}

// Method to store a string value into a key
func Put(key, value string) (Cache, error) {
	// Validate key
	if key == "" {
		return Cache{}, fmt.Errorf("'key' must not be nil")
	}

	// Validate value
	if value == "" {
		return Cache{}, fmt.Errorf("'value' must not be nil")
	}

	//  Create a Cache variable
	c := Cache{
		key:   key,
		value: []byte(value),
	}

	// Set cache path and filename
	c.SetPath()

	// Return Cache
	return c, nil
}

// Generate cache's path and filename
// Key is encrypted using SHA1
// First 2 characters are the parent folder name
// Third and Forth characters are the childrens folder name
// The entire hash is the filename
func (c *Cache) SetPath() {
	// Encrypt key and convert into byte array
	keyBytes := sha1.Sum([]byte(c.key))

	// Name parent and children folder
	c.Path = fmt.Sprintf("%x/%x/", keyBytes[0], keyBytes[1])

	// Loop to iterate the hash bytes
	// Generate filename's
	for i := 0; i < len(keyBytes); i++ {
		c.FileName += fmt.Sprintf("%x", keyBytes[i])
	}
}

// Save cache's value into a file
func (c *Cache) Save() error {
	// Dumping bytes into a file

	if _, err := os.Stat(c.Path); os.IsNotExist(err) {
		os.MkdirAll(c.Path, 0777)
	}

	f, err := os.Create(fmt.Sprintf("%s%s", c.Path, c.FileName))

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(c.value)

	if err != nil {
		return err
	}

	return nil
}

// Check if Cache already exists
func (c *Cache) Exists() (error, bool) {
	file, err := os.Stat(c.Path)

	if os.IsNotExist(err) {
		return fmt.Errorf("File does not exists"), false
	}
	isDir := !file.IsDir()

	// If file does not exists, if it's a directory
	if isDir {
		return fmt.Errorf("It's a directory not a file"), false
	}
	return nil, true
}
