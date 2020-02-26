package main

import (
	"testing"

	"github.com/EricNeid/go-getdockerimage/internal/verify"
)

func TestGetOutputName(t *testing.T) {
	// action
	result, err := getOutputName("foo")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo.docker.img", result)

	// action
	result, err = getOutputName("foo:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_2.0.0.docker.img", result)

	// action
	result, err = getOutputName("foo/bar")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_bar.docker.img", result)

	// action
	result, err = getOutputName("foo/bar:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_bar_2.0.0.docker.img", result)
}
