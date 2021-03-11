package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
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
func Store(t interface{}) (Cache, error) {
	c := Cache{}

	switch t.(type) {
	case stores.File:
		c.storeType = stores.File{}
		return c, nil
	case stores.Redis:
		c.storeType = stores.Redis{}
		return c, nil
	default:
		return Cache{}, fmt.Errorf("The format isn't supported")
	}
}

func (c Cache) Put(key, value string) error {
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
		return fmt.Errorf("The format isn't supported")
	}
}

func (c Cache) Get(key string) (string, error) {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Get(&FileCache)
	default:
		return "", fmt.Errorf("The format isn't supported")
	}
}

func (c Cache) Has(key string) (bool, error) {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Has(&FileCache)
	default:
		return false, fmt.Errorf("The format isn't supported")
	}
}

func (c Cache) Delete(key string) error {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Delete(&FileCache)
	default:
		return fmt.Errorf("The format isn't supported")
	}
}

func (c Cache) Pull(key string) (bool, error) {
	c.key = key

	switch c.storeType.(type) {
	case stores.File:
		FileCache := stores.File{
			Key: key,
		}
		return stores.Has(&FileCache)
	default:
		return false, fmt.Errorf("The format isn't supported")
	}
}
