package singleflight

import (
	"errors"
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	var g Group
	v, err := g.Do("key", func() (i interface{}, err error) {
		return "bar", nil
	})

	if got, want := fmt.Sprintf("%v (%T)", v, v), "bar (string)"; got != want {
		t.Errorf("Do = %v; want %v", got, want)
	}

	if err != nil {
		t.Errorf("Do error = %v", err)
	}
}

func TestDoErr(t *testing.T) {
	var g Group
	someErr := errors.New("some error")

	v, err := g.Do("key", func() (i interface{}, err error) {
		return nil, someErr
	})

	if err != someErr {
		t.Errorf("Do error = %v;want someErr", err)
	}

	if v != nil {
		t.Errorf("unexpected non-nil value %#v", v)
	}
}
