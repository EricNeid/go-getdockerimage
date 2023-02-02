// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT

package gogetdockerimage

import (
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// RemoteDestination represents a destination on a remote host.
// It contains user name, remote address (containing port number)
// and folder on remote host.
type RemoteDestination struct {
	User string
	Pass string
	Addr string
}

// ParseDestination parses the given url and returns new RemoteDestination.
// Url is expected to be in the format: ssh://user@10.20.300.400:22.
func ParseDestination(urlString string) (*RemoteDestination, error) {
	res, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	pass, _ := res.User.Password()
	return &RemoteDestination{
		User: res.User.Username(),
		Pass: pass,
		Addr: res.Host,
	}, nil
}

// SSHCopyFile copies a file from srcPath to dstPath on the provided host.
// Host url can include port (ie. "10.20.300.400:22")
func SSHCopyFile(user, pass, addr, srcPath, dstPath string) error {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			fmt.Println()
			return nil
		}, //ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return err
	}
	defer client.Close()

	// open an SFTP session over an existing ssh connection.
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	// Open the source file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := sftpClient.Create(dstPath)
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
