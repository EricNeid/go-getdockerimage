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

Checkout the project our simply click on the Gitpod link above.

```bash
make all
```

## Usage

If go is installed and $GOPATH/bin is in your path, you can download and install the tool directly
by using go tools.

Download tool:

```bash
go get github.com/EricNeid/go-getdockerimage/cmd/getdockerimage
go install github.com/EricNeid/go-getdockerimage/cmd/getdockerimage
```

Run it:

```bash
getdockerimage.exe foo/image:2.0.0
=> foo_image_2.0.0.img
```

## Question or comments

Please feel free to open a new issue:
<https://github.com/EricNeid/go-getdockerimage/issues>
