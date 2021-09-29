// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const version = "0.3.1"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Printf("Version: %s\n", version)
		flag.PrintDefaults()

		fmt.Printf("  Usage  : %s <image-name>\n", os.Args[0])
		fmt.Printf("  Example: %s foo/image:2.0.0\n", os.Args[0])
	}
	flag.Parse()

	if len(os.Args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	image := os.Args[1]
	output, err := getOutputName(image)
	if err != nil {
		fmt.Println("Error while generating output name " + err.Error())
		os.Exit(1)
	}

	err = downloadImage(image)
	if err != nil {
		fmt.Println("Error while downloading docker image " + err.Error())
		os.Exit(1)
	}

	saveImage(image, output)
}

func downloadImage(image string) error {
	fmt.Printf("docker pull %s\n", image)
	cmd := exec.Command("docker", "pull", image)

	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	go printBuffer(&outBuff)
	go printBuffer(&errBuff)

	return cmd.Run()
}

func saveImage(image string, output string) error {
	fmt.Printf("docker save %s --output %s\n", image, output)
	cmd := exec.Command("docker", "save", image, "--output", output)

	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	go printBuffer(&outBuff)
	go printBuffer(&errBuff)

	return cmd.Run()
}

// getOutputName returns name of generated file from image name
// foo -> foo.docker.img
// foo/bar -> foo_bar.docker.img
// foo/bar:42 -> foo_bar_42.docker.img
func getOutputName(image string) (string, error) {
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

func printBuffer(buffer *bytes.Buffer) {
	for {
		if buffer.Len() > 0 {
			next := buffer.Next(buffer.Len())
			fmt.Print(string(next))
		}
	}
}
