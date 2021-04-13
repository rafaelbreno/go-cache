package cache

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	redis "github.com/go-redis/redis/v8"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	cache "github.com/rafaelbreno/go-cache"
	"github.com/rafaelbreno/go-cache/pkg_error"
	"github.com/rafaelbreno/go-cache/stores"
)

var ctx = context.TODO()

type testRedis struct {
	name string
	want pkg_error.PkgError
	got  pkg_error.PkgError
}

// return a *redismock.ClientMock
func setRedisMocker() error {
	mr, err := miniredis.Run()
	if err != nil {
		// Can't run tests without this
		// So just panic
		return err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	stores.SetRedisConn(redisClient)

	return nil
}

// Prepare mocker
func Test_Mock_Redis(t *testing.T) {
	err := setRedisMocker()

	if err != nil {
		t.Errorf("\nWant: %v\nGot: %v", nil, err)
		// Closing
		// Without Redis mocker can't continue test
		panic(err)
	}
}

func getStoreRedisPutTests() []testRedis {
	var t []testRedis

	_, err := cache.Store(2)
	t = append(t, testRedis{
		name: "Incorrect Type",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.InvalidFormat, "int"),
		got: err,
	})

	redisCache, _ := cache.Store(stores.Redis{})
	err = redisCache.Put("", "bar")

	t = append(t, testRedis{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	err = redisCache.Put("foo", "")

	t = append(t, testRedis{
		name: "Value missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "value"),
		got: err,
	})

	err = redisCache.Put("foo", "bar")
	t = append(t, testRedis{
		name: "Cache successfully put",
		want: pkg_error.NewNilError(),
		got:  err,
	})
	return t
}

func Test_Redis_Put(t *testing.T) {
	tts := getStoreRedisPutTests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func getStoreRedisGetTests() []testRedis {
	var t []testRedis

	_, err := cache.Store(2)
	t = append(t, testRedis{
		name: "Incorrect Type",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.InvalidFormat, "int"),
		got: err,
	})

	redisCache, _ := cache.Store(stores.Redis{})
	_, err = redisCache.Get("")

	t = append(t, testRedis{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	f2Key := "random_key"
	_, err = redisCache.Get(f2Key)

	t = append(t, testRedis{
		name: "Value Found",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, f2Key),
		got: err,
	})

	_, err = redisCache.Get("bar")
	t = append(t, testRedis{
		name: "Value not found",
		want: pkg_error.NewNilError(),
		got:  err,
	})
	return t
}

func Test_Redis_Get(t *testing.T) {
	tts := getStoreRedisGetTests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func getStoreRedisHasTests() []testRedis {
	var t []testRedis

	redisCache, _ := cache.Store(stores.Redis{})
	_, err := redisCache.Has("")
	t = append(t, testRedis{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	_, err = redisCache.Has("foo")

	t = append(t, testRedis{
		name: "Value Found",
		want: pkg_error.NewNilError(),
		got:  err,
	})

	_, err = redisCache.Has("bar")
	t = append(t, testRedis{
		name: "Value not found",
		want: pkg_error.NewNilError().SetMessage(pkg_error.CacheDontExists, "bar"),
		got:  err,
	})
	return t
}

func Test_Redis_Has(t *testing.T) {
	tts := getStoreRedisGetTests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func getStoreRedisPullTests() []testRedis {
	var t []testRedis

	redisCache, _ := cache.Store(stores.Redis{})
	_, err := redisCache.Has("")
	t = append(t, testRedis{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	_, err = redisCache.Has("foo")

	t = append(t, testRedis{
		name: "Value Found",
		want: pkg_error.NewNilError(),
		got:  err,
	})

	_, err = redisCache.Has("bar")
	t = append(t, testRedis{
		name: "Value not found",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, "bar"),
		got: err,
	})
	return t
}

func Test_Redis_Pull(t *testing.T) {
	tts := getStoreRedisPullTests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func getStoreRedisDeleteTests() []testRedis {
	var t []testRedis

	redisCache, _ := cache.Store(stores.Redis{})
	err := redisCache.Delete("")
	t = append(t, testRedis{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, "key"),
		got: err,
	})

	err = redisCache.Delete("foo")

	t = append(t, testRedis{
		name: "Value Found",
		want: pkg_error.NewNilError(),
		got:  err,
	})

	err = redisCache.Delete("bar")
	t = append(t, testRedis{
		name: "Value not found",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, "bar"),
		got: err,
	})
	return t
}

func Test_Redis_Delete(t *testing.T) {
	tts := getStoreRedisDeleteTests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}
