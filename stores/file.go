package stores

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rafaelbreno/go-cache/pkg_error"
)

// TODO: implement expiration
type File struct {
	Key      string // Cache identifier - Cannot be null
	Value    []byte // Cache value itself
	path     string // Cache file path with the stored value
	fileName string // Cache file's name
}

// Method to store a string value into a key
func (f *File) Put() pkg_error.PkgError {
	// Validate key
	if f.Key == "" {
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	// Validate value
	if f.Value == nil {
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "value")
	}

	// Set cache path and filename
	f.SetPath()

	// Save value into
	return f.Save()
}

// Retrieve cached value
func (f *File) Get() (string, pkg_error.PkgError) {
	// Validate key
	if f.Key == "" {
		return "", pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	if has, err := f.Has(); !has {
		return "", err
	}

	dat, err := ioutil.ReadFile(f.fileName)
	if err != nil {
		return "", pkg_error.
			NewError(err)
	}

	return string(dat), pkg_error.NewNilError()
}

// Check if Cache already exists
func (f *File) Has() (bool, pkg_error.PkgError) {
	// Validate key
	if f.Key == "" {
		return false, pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	file, err := os.Stat(f.path)

	if os.IsNotExist(err) {
		return false, pkg_error.
			NewError(err).
			SetMessage(pkg_error.CacheDontExists, f.Key)
	}
	isDir := !file.IsDir()

	// If file does not exists, if it's a directory
	if isDir {
		return false, pkg_error.
			NewError(err).
			SetMessage(pkg_error.CacheDontExists, f.Key)
	}
	return true, pkg_error.NewNilError()
}

// Delete cached file
func (f *File) Delete() pkg_error.PkgError {
	// Validate key
	if f.Key == "" {
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	if err := os.Remove(f.path); err != nil {
		return pkg_error.
			NewError(err)
	}

	return pkg_error.NewNilError()
}

// Retrieve and delete cached file
func (f *File) Pull() (string, pkg_error.PkgError) {
	// Validate key
	if f.Key == "" {
		return "", pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	val, errPkg := f.Get()

	if errPkg.Nil {
		return "", errPkg
	}

	errPkg = f.Delete()

	if errPkg.Nil {
		return "", errPkg
	}

	return val, pkg_error.NewNilError()
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
func (f *File) Save() pkg_error.PkgError {
	// Dumping bytes into a file

	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		os.MkdirAll(f.path, 0777)
	}

	cacheFile, err := os.Create(fmt.Sprintf("%s%s", f.path, f.fileName))

	if err != nil {
		return pkg_error.
			NewError(err)
	}

	defer cacheFile.Close()

	_, err = cacheFile.Write(f.Value)

	if err != nil {
		return pkg_error.
			NewError(err)
	}

	return pkg_error.NewNilError()
}
