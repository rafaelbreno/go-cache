package cache_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	cache "github.com/rafaelbreno/go-cache"
	"github.com/rafaelbreno/go-cache/stores"
)

type testFile struct {
	name string
	want error
	got  error
}

func get_Store_File_Put_Tests() []testFile {
	var t []testFile

	_, err := cache.Store(2)

	t = append(t, testFile{
		name: "Incorrect Type",
		want: fmt.Errorf("The format isn't supported"),
		got:  err,
	})

	f2, _ := cache.Store(stores.File{})
	err = f2.Put("", "bar")

	t = append(t, testFile{
		name: "Key missing",
		want: fmt.Errorf("'key' must not be nil"),
		got:  err,
	})

	f3, _ := cache.Store(stores.File{})
	err = f3.Put("foo", "")

	t = append(t, testFile{
		name: "Value missing",
		want: fmt.Errorf("'value' must not be nil"),
		got:  err,
	})

	f4, _ := cache.Store(stores.File{})
	err = f4.Put("foo", "bar")

	t = append(t, testFile{
		name: "Cache successfully put",
		want: nil,
		got:  err,
	})

	return t
}

func Test_Store_File_Put(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Put_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != nil && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func get_Store_File_Has_Tests() []testFile {
	var t []testFile

	f1, _ := cache.Store(stores.File{})

	_, err := f1.Has("")

	t = append(t, testFile{
		name: "Key missing",
		want: fmt.Errorf("'key' must not be nil"),
		got:  err,
	})

	f2, _ := cache.Store(stores.File{})
	_, err = f2.Has("random_key")

	t = append(t, testFile{
		name: "Cache not found",
		want: fmt.Errorf("Cache doesn't exists"),
		got:  err,
	})

	f3, _ := cache.Store(stores.File{})
	_, err = f3.Has("foo")

	t = append(t, testFile{
		name: "Value exists",
		want: nil,
		got:  err,
	})

	return t
}

func Test_Store_File_Has(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Has_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != nil && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func get_Store_File_Get_Tests() []testFile {
	var t []testFile

	f1, _ := cache.Store(stores.File{})

	_, err := f1.Get("")

	t = append(t, testFile{
		name: "Key missing",
		want: fmt.Errorf("'key' must not be nil"),
		got:  err,
	})

	f2, _ := cache.Store(stores.File{})
	_, err = f2.Get("random_key")

	t = append(t, testFile{
		name: "Cache not found",
		want: fmt.Errorf("Cache doesn't exists"),
		got:  err,
	})

	f3, _ := cache.Store(stores.File{})
	_, err = f3.Get("foo")

	t = append(t, testFile{
		name: "Value exists",
		want: nil,
		got:  err,
	})

	return t
}

func Test_Store_File_Get(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Get_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != nil && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}
