# Bouquetin
Bouquetin is a file manager.
It will run on Windows, Linux natively, and Android via a bridge app.

## Run locally
Go will be required on the host to build for now (sad).

## Build
First, make sure gomobile is installed and in PATH:
```bash
go install golang.org/x/mobile/cmd/gomobile@latest # will most likely install in ~/go/bin
gomobile init
```

Make sure the exported android package is named `core`.
Make sure the `core` package is initialized (`cd core; go mod init core`).

Make sure you've installed Android NDK (can be done via Android Studio > tools > SDK manager > SDK tools > NDK)
Make sure you've installed javac.

Then:
```bash
cd bouquetin
./build.sh
```

This will output two binaries for W and Linux, and a .aar file to import in the Bouquetin Android bridge app.

(build for android under the hood):
go get golang.org/x/mobile/cmd/gobind
gomobile bind -v -o bqt.aar -target=android -androidapi 19 .

A docker image to build is not available for now.
