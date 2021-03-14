package cache

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	cache "github.com/rafaelbreno/go-cache"
	"github.com/rafaelbreno/go-cache/stores"
)

var ctx = context.TODO()

type testRedis struct {
	name string
	want error
	got  error
}

// return a *redismock.ClientMock
func init() {
	mr, err := miniredis.Run()
	if err != nil {
		// Can't run tests without this
		// So just panic
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	stores.SetRedisConn(redisClient)
}

func getStoreRedisPutTests() []testRedis {
	var t []testRedis

	_, err := cache.Store(2)
	t = append(t, testRedis{
		name: "Incorrect Type",
		want: fmt.Errorf("The format isn't supported"),
		got:  err,
	})

	redisCache, _ := cache.Store(stores.Redis{})
	err = redisCache.Put("", "bar")

	t = append(t, testRedis{
		name: "Key missing",
		want: fmt.Errorf("'key' must not be nil"),
		got:  err,
	})

	err = redisCache.Put("foo", "")

	t = append(t, testRedis{
		name: "Value missing",
		want: fmt.Errorf("'value' must not be nil"),
		got:  err,
	})

	err = redisCache.Put("foo", "bar")
	t = append(t, testRedis{
		name: "Cache successfully put",
		want: nil,
		got:  err,
	})
	return t
}

func Test_Redis_Put(t *testing.T) {
	tts := getStoreRedisPutTests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != nil {
				if tt.want.Error() != tt.got.Error() {
					t.Errorf("\nWant: %v\n Got: %v\n", tt.want, tt.got)
				}
			}
		})
	}
}
