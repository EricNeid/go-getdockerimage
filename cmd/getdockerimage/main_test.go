package main

import "testing"
import "EricNeid/go-getdockerimage/internal/verify"

func TestGetOutputName(t *testing.T) {
	// action
	result, err := getOutputName("foo")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo.img", result)

	// action
	result, err = getOutputName("foo:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_2.0.0.img", result)

	// action
	result, err = getOutputName("foo/bar")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_bar.img", result)

	// action
	result, err = getOutputName("foo/bar:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_bar_2.0.0.img", result)
}
