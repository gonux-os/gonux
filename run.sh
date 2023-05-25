#!/bin/sh

cd linux
qemu-system-x86_64 -kernel arch/x86/boot/bzImage -initrd "../rootfs.cpio.gz"
cd ..
