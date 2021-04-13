package cache_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	cache "github.com/rafaelbreno/go-cache"
	"github.com/rafaelbreno/go-cache/pkg_error"
	"github.com/rafaelbreno/go-cache/stores"
)

type testFile struct {
	name string
	want pkg_error.PkgError
	got  pkg_error.PkgError
}

func get_Store_File_Put_Tests() []testFile {
	var t []testFile

	_, err := cache.Store(2)

	t = append(t, testFile{
		name: "Incorrect Type",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.InvalidFormat, "int"),
		got: err,
	})

	f2, _ := cache.Store(stores.File{})
	err = f2.Put("", "bar")

	t = append(t, testFile{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	f3, _ := cache.Store(stores.File{})
	err = f3.Put("foo", "")

	t = append(t, testFile{
		name: "Value missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "value"),
		got: err,
	})

	f4, _ := cache.Store(stores.File{})
	err = f4.Put("foo", "bar")

	t = append(t, testFile{
		name: "Cache successfully put",
		want: pkg_error.NewNilError(),
		got:  err,
	})

	return t
}

func Test_Store_File_Put(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Put_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
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
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	f2, _ := cache.Store(stores.File{})
	f2Key := "random_key"
	_, err = f2.Has(f2Key)

	t = append(t, testFile{
		name: "Cache not found",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, f2Key),
		got: err,
	})

	f3, _ := cache.Store(stores.File{})
	_, err = f3.Has("foo")

	t = append(t, testFile{
		name: "Value exists",
		want: pkg_error.NewNilError(),
		got:  err,
	})

	return t
}

func Test_Store_File_Has(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Has_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
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
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	f2, _ := cache.Store(stores.File{})
	f2Key := "random_key"
	_, err = f2.Get(f2Key)

	t = append(t, testFile{
		name: "Cache not found",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, f2Key),
		got: err,
	})

	f3, _ := cache.Store(stores.File{})
	_, err = f3.Get("foo")

	t = append(t, testFile{
		name: "Value exists",
		want: pkg_error.NewNilError(),
		got:  err,
	})

	return t
}

func Test_Store_File_Get(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Get_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func get_Store_File_Pull_Tests() []testFile {
	var t []testFile

	f1, _ := cache.Store(stores.File{})

	_, err := f1.Pull("")

	t = append(t, testFile{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	f2, _ := cache.Store(stores.File{})
	f2Key := "random_key"
	_, err = f2.Pull(f2Key)

	t = append(t, testFile{
		name: "Cache not found",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, f2Key),
		got: err,
	})

	f3, _ := cache.Store(stores.File{})
	_, err = f3.Pull("foo")

	t = append(t, testFile{
		name: "Value retrieve and cache deleted",
		want: pkg_error.NewNilError(),
		got:  err,
	})
	return t
}

func Test_Store_File_Pull(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Get_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}

func get_Store_File_Delete_Tests() []testFile {
	// Inserting value
	f4, _ := cache.Store(stores.File{})
	_ = f4.Put("foo", "bar")

	var t []testFile

	f1, _ := cache.Store(stores.File{})

	err := f1.Delete("")

	t = append(t, testFile{
		name: "Key missing",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.FieldMustNotBeNull, "key"),
		got: err,
	})

	f2, _ := cache.Store(stores.File{})
	f2Key := "random_key"
	err = f2.Delete(f2Key)

	t = append(t, testFile{
		name: "Cache not found",
		want: pkg_error.
			NewNilError().
			SetMessage(pkg_error.CacheDontExists, f2Key),
		got: err,
	})

	f3, _ := cache.Store(stores.File{})
	err = f3.Delete("foo")

	t = append(t, testFile{
		name: "Value retrieve and cache deleted",
		want: pkg_error.NewNilError(),
		got:  err,
	})
	return t
}

func Test_Store_File_Delete(t *testing.T) {
	t.Helper()

	tts := get_Store_File_Delete_Tests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.IsNil() && cmp.Equal(tt.want, tt.got, cmpopts.EquateErrors()) {
				t.Errorf("\nWant: %v\nGot: %v", tt.want, tt.want)
			}
		})
	}
}
