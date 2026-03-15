#!/bin/bash

CGO_ENABLED=0 go build -o filesystem-tweaks/usr/local/bin/kstart cmd/*
fuse-ext2 ~/.local/share/uml/images/koble-fs /mnt/nktest -o rw+
sudo cp filesystem-tweaks/usr/local/bin/kstart /mnt/nktest/usr/local/bin/kstart
umount /mnt/nktest

