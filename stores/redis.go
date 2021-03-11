package stores

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"

	cache "github.com/rafaelbreno/go-cache"
	"github.com/go-redis/redis/v8"
)

// TODO: implement expiration
type Redis struct {
	Key      string // Cache identifier - Cannot be null
	Value    []byte // Cache value itself
	path     string // Cache file path with the stored value
	fileName string // Cache file's name
}

var redisClient *redis.Client

func init() {
	redisClient = cache.RedisClient
}

// Method to store a string value into a key
func (f *Redis) Put() error {
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
func (r *Redis) Get() (string, error) {
	// Validate key
	if r.Key == "" {
		return "", fmt.Errorf("'key' must not be nil")
	}

	var strCmd redis.StringCmd
	var has bool
	var err error

	if has, err, strCmd = r.Has(); !has {
		return "", err
	}

	return strCmd.String(), nil
}

// Check if Cache already exists
func (r *Redis) Has() (bool, error, redis.StringCmd) {
	// Validate key
	if r.Key == "" {
		return false, fmt.Errorf("'key' must not be nil"), redis.StringCmd{}
	}

	// Redis `GET key` command. It returns redis.Nil error when key does not exist.
	strCmd := redisClient.Get(r.Key)

	if err := strCmd.Err(); err != nil {
		return false, err, redis.StringCmd{}
	}

	return true, nil, *strCmd
}

// Delete cached file
func (r *Redis) Delete() error {
	// Validate key
	if r.Key == "" {
		return fmt.Errorf("'key' must not be nil")
	}

	err := redisClient.Del(r.Key).Err()

	return err
}

// Retrieve and delete cached file
func (r *Redis) Pull() (string, error) {
	// Validate key
	if r.Key == "" {
		return "", fmt.Errorf("'key' must not be nil")
	}

	val, err := r.Get()

	if err != nil {
		return "", err
	}

	err = r.Delete()

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
func (f *Redis) SetPath() {
	// Encrypt key and convert into byte array
	keyBytes := sha1.Sum([]byte(f.Key))

	// Loop to iterate the hash bytes
	// Generate filename's
	for i := 0; i < len(keyBytes); i++ {
		f.fileName += fmt.Sprintf("%x", keyBytes[i])
	}
}

// Save cache's value into a file
func (r *Redis) Save() error {
	// Dumping bytes into a file

	err := redisClient.Set(r.Key, r.fileName, 0).Err()

	return err
}
