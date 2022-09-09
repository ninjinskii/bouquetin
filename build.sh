#!/bin/bash

mkdir bin
mkdir bin/linux
mkdir bin/windows
mkdir bin/android

go build -o bin/linux/bqt
GOOS=windows GOARCH=amd64 go build -o bin/windows/bqt.exe

# cd core
# go get golang.org/x/mobile/cmd/gobind
# gomobile bind -v -o bqt.aar -target=android -androidapi 19 .
