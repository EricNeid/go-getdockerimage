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
		if len(parts) != 2 {
			return "", errors.New("Unexpected image name format " + image)
		}
		group = parts[0]
		img = parts[1]
		image = parts[1]
	}

	if strings.Contains(image, ":") {
		parts := strings.Split(image, ":")
		if len(parts) != 2 {
			return "", errors.New("Unexpected image name format " + image)
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

	output = output + ".docker.img"

	return output, nil
}

// DownloadImage download docker image by name.
func DownloadImage(imageName string) error {
	fmt.Printf("docker pull %s\n", imageName)
	cmd := exec.Command("docker", "pull", imageName)

	writer := io.MultiWriter(os.Stdout)
	cmd.Stdout = writer
	cmd.Stderr = writer

	return cmd.Run()
}

// SaveImage write docker image to file system.
func SaveImage(imageName, outDir, outName string) error {
	fmt.Printf("docker save %s --output %s\n", imageName, outName)
	cmd := exec.Command("docker", "save", imageName, "--output", outName)

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

// RemoveDir deletes directory and all files in it.
func RemoveDir(dir string) error {
	return os.RemoveAll(dir)
}
