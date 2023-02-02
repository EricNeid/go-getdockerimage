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

const (
	version     = "0.6.0"
	dockerfile  = "DOCKERFILE"
	composeFile = "DOCKER-COMPOSE.YML"
)

var (
	outDir = ""
	ssh    = ""
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Printf("Version: %s\n", version)
		flag.PrintDefaults()

		fmt.Printf("  Usage  : %s <image-name|dockerfile|docker-compose.yml|docker-project-folder>\n", os.Args[0])
		fmt.Printf("  Example: %s foo/image:2.0.0\n", os.Args[0])
	}
	flag.StringVar(&outDir, "dir", outDir, "(Optional) output directory")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	input := flag.Args()[0]
	f, err := os.Stat(input)

	switch {
	case errors.Is(err, os.ErrNotExist):
		// image name given
		handleImage(input)
	case err != nil:
		// could be custom registry
		handleImage(input)
	case f.IsDir():
		// directory given
		handleDir(input)
	case strings.ToUpper(filepath.Base(input)) == dockerfile:
		// dockerfile given
		handleDockerFile(input)
	case strings.ToUpper(filepath.Base(input)) == composeFile:
		// docker-compose file given
		handleDockerComposeFile(input)
	default:
		fmt.Println("Argument not understood, expecting image|dockerfile|docker-compose.yml|directory")
		fmt.Printf("Arguments was %s\n", input)
		os.Exit(1)
	}
}

func handleImage(image string) {
	fmt.Printf("Handling image %s\n", image)

	output, errOutputName := getdockerimage.GetOutputName(image)
	if errOutputName != nil {
		fmt.Println("Error while generating output name", errOutputName)
		os.Exit(1)
	}

	errDownload := getdockerimage.DownloadImage(image)
	if errDownload != nil {
		fmt.Println("Error while downloading docker image", errDownload)
		os.Exit(1)
	}

	errSaveImage := getdockerimage.SaveImage(image, outDir, output)
	if errSaveImage != nil {
		fmt.Println("Error while saving docker image", errSaveImage)
		os.Exit(1)
	}
}

func handleDockerFile(dockerFile string) {
	fmt.Printf("Handling dockerfile %s\n", dockerFile)
	images, err := getdockerimage.GetImagesFromDockerfile(dockerFile)
	if err != nil {
		fmt.Printf("Error while handling dockerfile %s %s\n", dockerFile, err)
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
		fmt.Printf("Error while handling docker-compose file %s %s\n", dockerComposeFile, err)
		os.Exit(1)
	}
	for _, image := range images {
		handleImage(image)
	}
}

func handleDir(dir string) {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error while walking directory %s %s\n", path, err)
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
		fmt.Printf("Error while handling directory %s %s\n", dir, err)
		os.Exit(1)
	}
}
