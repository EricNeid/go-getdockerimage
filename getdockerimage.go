// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT

// Package gogetdockerimage contains functions to download and save docker images.
package gogetdockerimage

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type DockerExecutable string

const PODMAN DockerExecutable = "podman"
const DOCKER DockerExecutable = "docker"

// GetCustomRegistry return the name of custom registry, if present
func GetCustomRegistry(image string) (string, error) {
	parts := strings.Split(image, "/")
	if len(parts) == 3 {
		return parts[0], nil
	}
	return "", errors.New("no registry was found")
}

// GetOutputName returns name of generated file from image name
// foo -> foo.docker.img
// foo/bar -> foo_bar.docker.img
// foo/bar:42 -> foo_bar_42.docker.img
func GetOutputName(image string) (string, error) {
	var group string
	var img string
	var version string

	img = image

	if strings.Contains(image, "/") {
		parts := strings.Split(image, "/")
		switch len(parts) {
		case 2:
			group = parts[0]
			img = parts[1]
			image = parts[1]
		case 3:
			// parts[0] is custom registry
			group = parts[1]
			img = parts[2]
			image = parts[2]
		default:
			return "", fmt.Errorf("unexpected image name format: %s ", image)
		}
	}

	if strings.Contains(image, ":") {
		parts := strings.Split(image, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("unexpected image name format: %s ", image)
		}
		img = parts[0]
		version = parts[1]
	}

	output := img

	if group != "" {
		output = group + "_" + output
	}

	if version != "" {
		output = output + "_" + version
	}

	output += ".docker.img"

	return output, nil
}

// GetDockerExecutable return the name of the docker executable (podman or docker).
func GetDockerExecutable() (DockerExecutable, error) {
	_, err := exec.LookPath("podman")
	if err == nil {
		return PODMAN, nil
	}
	_, err = exec.LookPath("docker")
	if err == nil {
		return DOCKER, nil
	}
	return "", errors.New("neither docker nor podman executable found")
}

// DownloadImage download docker image by name.
func DownloadImage(dockerExec DockerExecutable, imageName string) error {
	fmt.Printf("%s pull %s\n", dockerExec, imageName)

	// set command explicitly to prevent tainted input
	var cmd *exec.Cmd
	switch dockerExec {
	case PODMAN:
		cmd = exec.Command("podman", "pull", imageName)
	case DOCKER:
		cmd = exec.Command("docker", "pull", imageName)
	default:
		return errors.New("unsupported docker executable")
	}
	writer := io.MultiWriter(os.Stdout)
	cmd.Stdout = writer
	cmd.Stderr = writer

	return cmd.Run()
}

// SaveImage write docker image to file system.
func SaveImage(dockerExec DockerExecutable, imageName, outDir, outName string) error {
	fmt.Printf("%s save %s --output %s\n", dockerExec, imageName, outName)

	// set command explicitly to prevent tainted input
	var cmd *exec.Cmd
	switch dockerExec {
	case PODMAN:
		cmd = exec.Command("podman", "save", imageName, "--output", outName)
	case DOCKER:
		cmd = exec.Command("docker", "save", imageName, "--output", outName)
	default:
		return errors.New("unsupported docker executable")
	}
	writer := io.MultiWriter(os.Stdout)
	cmd.Stdout = writer
	cmd.Stderr = writer

	err := cmd.Run()
	if err != nil {
		return err
	}

	if outDir != "" {
		fmt.Printf("Moving %s to dir %s\n", outName, outDir)
		err = os.MkdirAll(outDir, os.ModePerm)
		if err != nil {
			return err
		}
		return os.Rename(outName, outDir+"/"+outName)
	}

	return nil
}
