package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage  : %s <image-name>\n", os.Args[0])
		fmt.Printf("Example: %s foo/image:2.0.0\n", os.Args[0])
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
	cmd := exec.Command("docker", "pull", image)

	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	err := cmd.Run()
	if err != nil {
		println(errBuff.String())
	} else {
		println(outBuff.String())
	}

	return err
}

func saveImage(image string, output string) error {
	cmd := exec.Command("docker", "save", image, "--output", output)

	var outBuff, errBuff bytes.Buffer
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	err := cmd.Run()
	if err != nil {
		println(errBuff.String())
	} else {
		println(outBuff.String())
	}

	return err
}

// getOutputName returns name of generated file from image name
// foo -> foo.img
// foo/bar -> foo_bar.img
// foo/bar:42 -> foo_bar_42.img
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

	output = output + ".img"

	return output, nil
}
