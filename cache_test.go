package cache

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_Without_Key(t *testing.T) {
	t.Helper()

	want := fmt.Errorf("'key' must not be nil")
	//want := fmt.Errorf("'value' must not be nil")
	if _, gotErr := Put("", "Go Cache"); cmp.Equal(want, gotErr, cmpopts.EquateErrors()) {
		t.Errorf("\nWant: %v\nGot: %v", want, gotErr)
	}

}

func Test_Without_Value(t *testing.T) {
	t.Helper()

	want := fmt.Errorf("'value' must not be nil")
	if _, gotErr := Put("name", ""); cmp.Equal(want, gotErr, cmpopts.EquateErrors()) {
		t.Errorf("\nWant: %v\nGot: %v, %v", want, gotErr, want == gotErr)
	}

}

func Test_Put(t *testing.T) {
	t.Helper()

	want := "/6a/e9/6ae999552ad2dca14d62e2bc8b764d377b1dd6c"
	if got, _ := Put("name", "Go Cache"); got.Path != want {
		t.Errorf("\nWant: %v\nGot:  %v", want, got.Path)
	}

}
