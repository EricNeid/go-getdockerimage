// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package gogetdockerimage

import (
	"os"
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

func TestRemoveDir(t *testing.T) {
	// arrange
	err := os.MkdirAll("testdata/tmp", os.ModePerm)
	verify.Ok(t, err)
	f, err := os.CreateTemp("testdata/tmp", "tmp.txt")
	verify.Ok(t, err)
	f.Close()
	// action
	err = RemoveDir("testdata/tmp")
	// verify
	verify.Ok(t, err)
	_, err = os.Stat("testdata/tmp")
	verify.Assert(t, os.IsNotExist(err), "directory not deleted")
}

func TestGetCustomRegistry(t *testing.T) {
	// action
	var result, err = GetCustomRegistry("myregistry.local:5000/foo/bar:2.0.0")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "myregistry.local:5000", result)
}
