// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	getdockerimage "github.com/EricNeid/go-getdockerimage"
)

const version = "0.4.1"
const dockerfile = "DOCKERFILE"
const composeFile = "DOCKER-COMPOSE.YML"

var destination = ""

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Printf("Version: %s\n", version)
		flag.PrintDefaults()

		fmt.Printf("  Usage  : %s <image-name|dockerfile|docker-compose.yml|docker-project-folder>\n", os.Args[0])
		fmt.Printf("  Example: %s foo/image:2.0.0\n", os.Args[0])
	}

	flag.StringVar(&destination, "dst", destination, "(Optional) destination directory")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	input := flag.Args()[0]
	f, err := os.Stat(input)

	if errors.Is(err, os.ErrNotExist) {
		// image name given
		handleImage(input)
	} else if f.IsDir() {
		// directory given
		handleDir(input)
	} else if strings.ToUpper(filepath.Base(input)) == dockerfile {
		// dockerfile given
		handleDockerFile(input)
	} else if strings.ToUpper(filepath.Base(input)) == composeFile {
		// docker-compose file given
		handleDockerComposeFile(input)
	} else {
		fmt.Println("Argument not understood, expecting image|dockerfile|docker-compose.yml|directory")
		fmt.Printf("Arguments was %s\n", input)
		os.Exit(1)
	}
}

func handleImage(image string) {
	fmt.Printf("Handling image %s\n", image)

	output, err := getdockerimage.GetOutputName(image)
	if err != nil {
		fmt.Println("Error while generating output name " + err.Error())
		os.Exit(1)
	}

	err = getdockerimage.DownloadImage(image)
	if err != nil {
		fmt.Println("Error while downloading docker image " + err.Error())
		os.Exit(1)
	}

	// TODO
	// if destination is remote server then save image to tmp directory
	// move to destination server
	// delete tmp directory
	err = getdockerimage.SaveImage(image, destination, output)
	if err != nil {
		fmt.Println("Error while saving docker image " + err.Error())
		os.Exit(1)
	}
}

func handleDockerFile(dockerFile string) {
	fmt.Printf("Handling dockerfile %s\n", dockerFile)
	images, err := getdockerimage.GetImagesFromDockerfile(dockerFile)
	if err != nil {
		fmt.Printf("Error while handling dockerfile %s %s\n", dockerFile, err.Error())
		os.Exit(1)
	}
	for _, image := range images {
		handleImage(image)
	}
}

func handleDockerComposeFile(dockerComposeFile string) {
	fmt.Printf("Handling compose file %s\n", dockerComposeFile)
	images, err := getdockerimage.GetImagesFromDockerCompose(dockerComposeFile)
	if err != nil {
		fmt.Printf("Error while handling docker-compose file %s %s\n", dockerComposeFile, err.Error())
		os.Exit(1)
	}
	for _, image := range images {
		handleImage(image)
	}
}

func handleDir(dir string) {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error while walking directory %s %s\n", path, err.Error())
			return nil
		}
		if !d.IsDir() {
			if strings.ToUpper(d.Name()) == dockerfile {
				handleDockerFile(path)
			}
			if strings.ToUpper(d.Name()) == composeFile {
				handleDockerComposeFile(path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error while handling directory %s %s\n", dir, err.Error())
		os.Exit(1)
	}
}
