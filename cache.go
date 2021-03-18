package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/rafaelbreno/go-cache/helpers"
	"github.com/rafaelbreno/go-cache/pkg_error"
	"github.com/rafaelbreno/go-cache/stores"
)

type Cache struct {
	storeType interface{} // Select which cache provider
	key       string      // Cache identifier - Cannot be null
	value     []byte      // Cache value itself
}

var RedisClient *redis.Client

func SetConfig(conns ...interface{}) {
	for _, conn := range conns {
		switch conn.(type) {
		case *redis.Client:
			RedisClient = conn.(*redis.Client)
			break
		default:
			break
		}
	}

}

// Choose a store type(file, redis, memcached, dynamodb, etc.)
func Store(t interface{}) (Cache, pkg_error.PkgError) {
	c := Cache{}

	switch t.(type) {
	case stores.File:
		c.storeType = stores.File{}
		return c, pkg_error.NewNilError()
	case stores.Redis:
		c.storeType = stores.Redis{}
		return c, pkg_error.NewNilError()
	default:
		return Cache{}, pkg_error.
			NewError(nil).
			SetMessage(pkg_error.InvalidFormat, helpers.GetType(c.storeType))
	}
}

func (c Cache) Put(key, value string) pkg_error.PkgError {
	c.key = key
	c.value = []byte(value)

	switch c.storeType.(type) {
	case stores.File:
		filecache := stores.File{
			Key:   key,
			Value: []byte(value),
		}
		return stores.Put(&filecache)
	case stores.Redis:
		filecache := stores.Redis{
			Key:   key,
			Value: []byte(value),
		}
		return stores.Put(&filecache)
	default:
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.InvalidFormat, helpers.GetType(c.storeType))
	}
}

func (c Cache) Get(key string) (string, pkg_error.PkgError) {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Get(&FileCache)
	case stores.Redis:
		RedisCache := stores.Redis{
			Key: key,
		}
		return stores.Get(&RedisCache)
	default:
		return "", pkg_error.
			NewError(nil).
			SetMessage(pkg_error.InvalidFormat, helpers.GetType(c.storeType))
	}
}

func (c Cache) Has(key string) (bool, pkg_error.PkgError) {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Has(&FileCache)
	case stores.Redis:
		RedisCache := stores.Redis{
			Key: key,
		}
		return stores.Has(&RedisCache)
	default:
		return false, pkg_error.
			NewError(nil).
			SetMessage(pkg_error.InvalidFormat, helpers.GetType(c.storeType))
	}
}

func (c Cache) Delete(key string) pkg_error.PkgError {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Delete(&FileCache)
	case stores.Redis:
		RedisCache := stores.Redis{
			Key: key,
		}
		return stores.Delete(&RedisCache)
	default:
		return pkg_error.
			NewError(nil).
			SetMessage(pkg_error.InvalidFormat, helpers.GetType(c.storeType))
	}
}

func (c Cache) Pull(key string) (string, pkg_error.PkgError) {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Pull(&FileCache)
	case stores.Redis:
		RedisCache := stores.Redis{
			Key: key,
		}
		return stores.Pull(&RedisCache)
	default:
		return "", pkg_error.
			NewError(nil).
			SetMessage(pkg_error.InvalidFormat, helpers.GetType(c.storeType))
	}
}
