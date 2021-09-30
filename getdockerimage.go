// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package gogetdockerimage

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// getOutputName returns name of generated file from image name
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

func DownloadImage(image string) error {
	fmt.Printf("docker pull %s\n", image)
	cmd := exec.Command("docker", "pull", image)

	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	go printBuffer(&outBuff)
	go printBuffer(&errBuff)

	return cmd.Run()
}

func SaveImage(image string, output string) error {
	fmt.Printf("docker save %s --output %s\n", image, output)
	cmd := exec.Command("docker", "save", image, "--output", output)

	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	go printBuffer(&outBuff)
	go printBuffer(&errBuff)

	return cmd.Run()
}

func printBuffer(buffer *bytes.Buffer) {
	for {
		if buffer.Len() > 0 {
			next := buffer.Next(buffer.Len())
			fmt.Print(string(next))
		}
	}
}
