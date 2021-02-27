package cache

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	cache "github.com/rafaelbreno/go-cache"
)

func Test_Without_Key(t *testing.T) {
	t.Helper()

	want := fmt.Errorf("'key' must not be nil")
	if _, gotErr := cache.Put("", "Go Cache"); cmp.Equal(want, gotErr, cmpopts.EquateErrors()) {
		t.Errorf("\nWant: %v\nGot: %v", want, gotErr)
	}

}

func Test_Without_Value(t *testing.T) {
	t.Helper()

	want := fmt.Errorf("'value' must not be nil")
	if _, gotErr := cache.Put("name", ""); cmp.Equal(want, gotErr, cmpopts.EquateErrors()) {
		t.Errorf("\nWant: %v\nGot: %v, %v", want, gotErr, want == gotErr)
	}
}

func Test_Put(t *testing.T) {
	t.Helper()

	want := "6a/e9/"
	if got, _ := cache.Put("name", "Go Cache"); got.Path != want {
		t.Errorf("\nWant: %v\nGot:  %v", want, got.Path)
	}
}

func Test_Save(t *testing.T) {
	t.Helper()

	c, err := cache.Put("name", "Go Cache")

	if err != nil {
		t.Errorf("\nWant: %v\nGot:  %v", nil, err.Error())
	}

	if err := c.Save(); err != nil {
		t.Errorf("\nWant: %v\nGot:  %v", nil, err.Error())
	}
}

func Test_Save_Already_Exists(t *testing.T) {
	t.Helper()

	c, err := cache.Put("name_test", "Go Cache")

	if err != nil {
		t.Errorf("\nWant: %v\nGot:  %v", nil, err.Error())
	}

	if err := c.Save(); err != nil {
		t.Errorf("\nWant: %v\nGot:  %v", nil, err.Error())
	}

	err, exists := c.Exists()
	if !exists {
		t.Errorf("\nWant: %v\nGot:  %v", true, exists)
	}

	want := fmt.Errorf("File does not exists")
	if err != nil {
		t.Errorf("\nWant: %v\nGot:  %v", want.Error(), err.Error())
	}
}
