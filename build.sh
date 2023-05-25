#!/bin/sh

mkdir -p d
go build -o d/init gonux/main.go
cd d
find . | cpio -o -H newc | gzip > ../rootfs.cpio.gz
cd ..
