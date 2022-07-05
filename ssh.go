// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT

package gogetdockerimage

import (
	"os"

	"net/url"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// RemoteDestination represents a destination on a remote host.
// It contains user name, remote address (containing port number)
// and folder on remote host.
type RemoteDestination struct {
	User    string
	Addr    string
	DstPath string
}

// ParseDestination checks if the given destination is an url
// to a remote host (ie. ssh://user@10.20.300.400:22/home/user/dir).
// If that is the case it returns new RemoteDestination from url
// components.
func ParseDestination(dst string) (*RemoteDestination, error) {
	_, err := url.Parse(dst)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// SSHCopyFile copies a file from srcPath to dstPath on the provided host.
// Host url can include port (ie. "10.20.300.400:22")
func SSHCopyFile(user, pass, addr, srcPath, dstPath string) error {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return err
	}
	defer client.Close()

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		return err
	}
	defer sftp.Close()

	// Open the source file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := sftp.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// write to file
	if _, err := dstFile.ReadFrom(srcFile); err != nil {
		return err
	}
	return nil
}
