// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package gogetdockerimage

import (
	"testing"

	"github.com/EricNeid/go-getdockerimage/internal/verify"
)

func TestGetOutputName(t *testing.T) {
	// action
	result, err := GetOutputName("foo")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo.docker.img", result)

	// action
	result, err = GetOutputName("foo:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_2.0.0.docker.img", result)

	// action
	result, err = GetOutputName("foo/bar")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_bar.docker.img", result)

	// action
	result, err = GetOutputName("foo/bar:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_bar_2.0.0.docker.img", result)

	// action
	result, err = GetOutputName("/foo/bar:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "foo_bar_2.0.0.docker.img", result)
}

func TestGetCustomRegistry(t *testing.T) {
	// action
	var result, err = GetCustomRegistry("myregistry.local:5000/foo/bar:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "myregistry.local:5000", result)
}

func TestGetDockerExecutable(t *testing.T) {
	// action
	result, err := GetDockerExecutable()
	// verify
	verify.Assert(t, result == DOCKER || result == PODMAN || err != nil, "result must be either DOCKER or PODMAN, or an error must be returned")
}
