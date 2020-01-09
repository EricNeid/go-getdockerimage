# go-getdockerimage

A utility tool to download images from docker hub and export them to image file.

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
