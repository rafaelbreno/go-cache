package stores

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

// TODO: implement expiration
type File struct {
	Key      string // Cache identifier - Cannot be null
	Value    []byte // Cache value itself
	path     string // Cache file path with the stored value
	fileName string // Cache file's name
}

// Method to store a string value into a key
func (f *File) Put() error {
	// Validate key
	if f.Key == "" {
		return fmt.Errorf("'key' must not be nil")
	}

	// Validate value
	if f.Value == nil {
		return fmt.Errorf("'value' must not be nil")
	}

	// Set cache path and filename
	f.SetPath()

	// Save value into
	return f.Save()
}

// Retrieve cached value
func (f *File) Get() (string, error) {
	// Validate key
	if f.Key == "" {
		return "", fmt.Errorf("'key' must not be nil")
	}

	if has, err := f.Has(); !has {
		return "", err
	}

	dat, err := ioutil.ReadFile(f.fileName)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

// Check if Cache already exists
func (f *File) Has() (bool, error) {
	// Validate key
	if f.Key == "" {
		return false, fmt.Errorf("'key' must not be nil")
	}

	file, err := os.Stat(f.path)

	if os.IsNotExist(err) {
		return false, fmt.Errorf("Cache doesn't exists")
	}
	isDir := !file.IsDir()

	// If file does not exists, if it's a directory
	if isDir {
		return false, fmt.Errorf("Cache file doesn't exists")
	}
	return true, nil
}

// Delete cached file
func (f *File) Delete() error {
	// Validate key
	if f.Key == "" {
		return fmt.Errorf("'key' must not be nil")
	}

	if err := os.Remove(f.path); err != nil {
		return err
	}

	return nil
}

// Retrieve and delete cached file
func (f *File) Pull() (string, error) {
	// Validate key
	if f.Key == "" {
		return "", fmt.Errorf("'key' must not be nil")
	}

	val, err := f.Get()

	if err != nil {
		return "", err
	}

	err = f.Delete()

	if err != nil {
		return "", err
	}

	return val, nil
}

// Generate cache's path and filename
// Key is encrypted using SHA1
// First 2 characters are the parent folder name
// Third and Forth characters are the childrens folder name
// The entire hash is the filename
func (f *File) SetPath() {
	// Encrypt key and convert into byte array
	keyBytes := sha1.Sum([]byte(f.Key))

	// Name parent and children folder
	f.path = fmt.Sprintf("%x/%x/", keyBytes[0], keyBytes[1])

	// Loop to iterate the hash bytes
	// Generate filename's
	for i := 0; i < len(keyBytes); i++ {
		f.fileName += fmt.Sprintf("%x", keyBytes[i])
	}
}

// Save cache's value into a file
func (f *File) Save() error {
	// Dumping bytes into a file

	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		os.MkdirAll(f.path, 0777)
	}

	cacheFile, err := os.Create(fmt.Sprintf("%s%s", f.path, f.fileName))

	if err != nil {
		return err
	}

	defer cacheFile.Close()

	_, err = cacheFile.Write(f.Value)

	if err != nil {
		return err
	}

	return nil
}
