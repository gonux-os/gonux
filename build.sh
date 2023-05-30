#!/bin/sh

sudo rm out/* -rf

echo "Making initramfs filesystem"

mkdir -p out/rootfs

# Enter out/rootfs
cd out/rootfs

mkdir -p bin
mkdir -p dev
# mkdir -p etc
# mkdir -p home
# mkdir -p mnt
mkdir -p proc
mkdir -p sys
# mkdir -p usr

cd dev
sudo MAKEDEV std console
cd ..

# Leave out/rootfs
cd ../..

echo "Building init"
go build -o out/rootfs/init gonux/init

echo "Building programs:"

echo "  - pwd"
go build -o out/rootfs/bin/pwd gonux/pwd

echo "  - ls"
go build -o out/rootfs/bin/ls gonux/ls

echo "  - clear"
go build -o out/rootfs/bin/clear gonux/clear

echo "  - gorgon"
go build -o out/rootfs/bin/gorgon gonux/gorgon

echo "Packaging rootfs"
cd out/rootfs
find . | cpio -o -H newc | gzip > ../rootfs.cpio.gz
cd ..

echo "Done"
