# Postar build script
# Author: fishgoddess
VERSION=v0.2.0-alpha
BUILD_TARGET=target

# Before building
echo "Start building..."
mkdir $BUILD_TARGET
echo "Building target directory: $BUILD_TARGET"

# Go build: windows, linux and darwin
echo "Building windows version..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $BUILD_TARGET/postar-$VERSION-windows.exe

echo "Building linux version..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $BUILD_TARGET/postar-$VERSION-linux
chmod +x $BUILD_TARGET/postar-$VERSION-linux

echo "Building darwin version..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $BUILD_TARGET/postar-$VERSION-darwin
chmod +x $BUILD_TARGET/postar-$VERSION-darwin

# After building
echo "Done."
