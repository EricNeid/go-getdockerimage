# About

Getdockerimage is a utility tool to download images from docker hub and export them.
Exporting images is useful, if you have a server without access to docker hub.

## Use case

* Export image on a machine with access
* Copy image to your server
* Load image with docker load --input MyImage

## Getting started

Download tool.

```bash
go install EricNeid/go-getdockerimage/cmd/getdockerimage
```

Make sure that $GOPATH/bin is in your path.

```bash
getdockerimage.exe foo/image:2.0.0
=> foo_image_2.0.0.img
```

## Question or comments

Please feel free to open a new issue:
<https://github.com/EricNeid/go-getdockerimage/issues>
