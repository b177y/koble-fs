#!/bin/bash

CGO_ENABLED=0 go build -o filesystem-tweaks/usr/local/bin/kstart cmd/*
wc=$(buildah from koble-deb-test)
buildah copy $wc filesystem-tweaks/usr/local/bin/kstart /usr/local/bin/kstart
buildah commit $wc koble-deb-test
podman tag localhost/koble-deb-test docker.io/b177y/koble-deb
