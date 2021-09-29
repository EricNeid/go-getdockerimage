<!--
SPDX-FileCopyrightText: 2021 Eric Neidhardt
SPDX-License-Identifier: CC-BY-4.0
-->
<!-- markdownlint-disable MD022 MD032 MD024-->
<!-- markdownlint-disable MD041-->
[![Go Report Card](https://goreportcard.com/badge/github.com/EricNeid/go-getdockerimage?style=flat-square)](https://goreportcard.com/report/github.com/EricNeid/go-getdockerimage)
![Go](https://github.com/EricNeid/go-sleep/workflows/Go/badge.svg)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/EricNeid/go-getdockerimage)
[![Release](https://img.shields.io/github/release/EricNeid/go-getdockerimage.svg?style=flat-square)](https://github.com/EricNeid/go-sleep/releases/latest)
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/EricNeid/go-getdockerimage)

# About

Go-getdockerimage is a utility tool to download images from docker hub and export them.
Exporting images is useful, if you have a server without access to docker hub.

## Use case

* Export image on a machine with access to docker hub
* Copy image to your server
* Load image with docker load --input MyImage

# Requirements

Docker is installed on your machine.

## Quickstart

Checkout the project and run make (given that go build chain is installed. 
Hint: You can simply click on the Gitpod link above and compile it online without setting up a tool chain.

```bash
make all
```

## Installation

If go is installed and $GOPATH/bin is in your path, you can download and install the tool directly
by using go tools.

Download tool:

```bash
go get github.com/EricNeid/go-getdockerimage/cmd/getdockerimage
```

Run it:

```bash
getdockerimage.exe foo/image:2.0.0
=> foo_image_2.0.0.img
```

## Question or comments

Please feel free to open a new issue:
<https://github.com/EricNeid/go-getdockerimage/issues>
