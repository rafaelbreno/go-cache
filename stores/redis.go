package stores

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rafaelbreno/go-cache/pkg_error"
)

// TODO: implement expiration
type Redis struct {
	Key      string // Cache identifier - Cannot be null
	Value    []byte // Cache value itself
	path     string // Cache file path with the stored value
	fileName string // Cache file's name
	strCMD   *redis.StringCmd
}

var redisClient *redis.Client

var ctx = context.TODO()

func SetRedisConn(r *redis.Client) {
	redisClient = r
}

// Method to store a string value into a key
func (f *Redis) Put() pkg_error.PkgError {
	// Validate key
	if f.Key == "" {
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	// Validate value
	if len(f.Value) == 0 {
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "value")
	}

	return pkg_error.NewNilError()
}

// Retrieve cached value
func (r *Redis) Get() (string, pkg_error.PkgError) {
	// Validate key
	if r.Key == "" {
		return "", pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	if has, err := r.Has(); !has {
		return "", err
	}

	return r.strCMD.String(), pkg_error.NewNilError()
}

// Check if Cache already exists
func (r *Redis) Has() (bool, pkg_error.PkgError) {
	// Validate key
	if r.Key == "" {
		return false, pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	// Redis `GET key` command. It returns redis.Nil pkg_error.PkgError when key does not exist.
	strCmd := redisClient.Get(ctx, r.Key)

	if err := strCmd.Err(); err != redis.Nil {
		return false, pkg_error.
			NewError(err).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	r.strCMD = strCmd

	return true, pkg_error.NewNilError()
}

// Delete cached file
func (r *Redis) Delete() pkg_error.PkgError {
	// Validate key
	if r.Key == "" {
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	err := redisClient.Del(ctx, r.Key).Err()

	return pkg_error.
		NewError(err)
}

// Retrieve and delete cached file
func (r *Redis) Pull() (string, pkg_error.PkgError) {
	// Validate key
	if r.Key == "" {
		return "", pkg_error.
			NewError(nil).
			SetMessage(pkg_error.FieldMustNotBeNull, "key")
	}

	val, err := r.Get()

	if !err.Nil {
		return "", err
	}

	err = r.Delete()

	if !err.Nil {
		return "", err
	}

	return val, pkg_error.NewNilError()
}

// Save cache's value into a file
func (r *Redis) Save() pkg_error.PkgError {
	// Dumping bytes into a file

	err := redisClient.
		Set(ctx, r.Key, r.Value, 0).
		Err()

	return pkg_error.
		NewError(err)
}
