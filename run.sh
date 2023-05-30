#!/bin/sh

echo "Running qemu"
qemu-system-x86_64 -kernel linux/arch/x86/boot/bzImage -initrd "out/rootfs.cpio.gz"
