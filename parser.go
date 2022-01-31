// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package gogetdockerimage

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

// GetImagesFromDockerfile returns all images found in given Dockerfile.
// It search for occurrences of "FROM imageName" or "FROM imageName AS alias"
// Search is performed case insensitive.
// Declaration of scratch is ignored.
func GetImagesFromDockerfile(dockerfilePath string) ([]string, error) {
	lines, err := readFileLines(dockerfilePath)
	if err != nil {
		return nil, err
	}

	var images []string

	for _, l := range lines {
		// search for possible image
		// line is comment
		if strings.HasPrefix(l, "#") {
			continue
		}

		// expected format is FROM imageName AS alias
		index := strings.Index(strings.ToUpper(l), "FROM ")
		// line does not contain imageName
		if index == -1 {
			continue
		}
		imageString := l[index+len("FROM "):]
		imageName := ""
		for _, r := range imageString {
			// reached end of imageName => stop search
			if unicode.IsSpace(r) && len(imageName) != 0 {
				break
			}
			// looking for start of imageName
			if unicode.IsSpace(r) {
				continue
			}
			imageName += string(r)
		}
		// if image is not Scratch => append to result
		if !strings.Contains(strings.ToUpper(imageName), "SCRATCH") {
			images = append(images, imageName)
		}
	}
	return images, nil
}

// GetImagesFromDockeCompose returns all images found in given docker-compose file.
// It search for occurences of "image: <imageName>" or "image: "<imageName>""
// Search is performed case insensitive.
func GetImagesFromDockerCompose(dockerfilePath string) ([]string, error) {
	lines, err := readFileLines(dockerfilePath)
	if err != nil {
		return nil, err
	}

	var images []string

	for _, l := range lines {
		// search for possible image
		// expected format is image: imageName
		index := strings.Index(strings.ToUpper(l), "IMAGE: ")
		// line does not contain imageName
		if index == -1 {
			continue
		}
		imageString := l[index+len("IMAGE: "):]
		imageName := ""
		for _, r := range imageString {
			// reached end of imageName => stop search
			if (unicode.IsSpace(r) || r == '"') && len(imageName) != 0 {
				break
			}
			// looking for start of imageName
			if unicode.IsSpace(r) || r == '"' {
				continue
			}
			imageName += string(r)
		}
		images = append(images, imageName)
	}
	return images, nil
}

func readFileLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}
