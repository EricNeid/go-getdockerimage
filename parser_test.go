// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package gogetdockerimage

import (
	"testing"

	"github.com/EricNeid/go-getdockerimage/internal/verify"
)

func TestGetImagesFromDockerfile_fileNotFound(t *testing.T) {
	// action
	_, err := GetImagesFromDockerfile("./testdata/no_file")
	// verify
	verify.NotNil(t, err, "Should return error not found")
}

func TestGetImagesFromDockerfile_shouldFindBuilder_shouldIgnoreScratch(t *testing.T) {
	// action
	result, err := GetImagesFromDockerfile("./testdata/Dockerfile")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 1, len(result))
	verify.Equals(t, "golang:1.13.0-alpine3.10", result[0])
}

func TestGetImagesFromDockerfile2_shouldIgnoreComments(t *testing.T) {
	// action
	result, err := GetImagesFromDockerfile("./testdata/Dockerfile_2")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 1, len(result))
	verify.Equals(t, "golang:1.13.0-alpine3.10", result[0])
}

func TestGetImagesFromDockerCompose_shouldFindImages(t *testing.T) {
	// action
	result, err := GetImagesFromDockerCompose("./testdata/docker-compose.yml")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, 2, len(result))
	verify.Equals(t, "kartoza/postgis:13-3.1", result[0])
	verify.Equals(t, "nginx", result[1])
}
