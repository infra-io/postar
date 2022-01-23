#!/bin/bash

VERSION=v0.3.0-alpha
echo "VERSION: $VERSION"
echo "----------------------------------------------------------------------"

# Check
WORKDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
echo "WORKDIR: $WORKDIR"
TARGET=$WORKDIR/target
echo "TARGET: $TARGET"
CONFIG_FILE=$WORKDIR/_examples/config/postar.ini
echo "CONFIG_FILE: $CONFIG_FILE"
LICENSE_FILE=$WORKDIR/LICENSE
echo "LICENSE_FILE: $LICENSE_FILE"
echo "----------------------------------------------------------------------"

echo "Please check the paths above! It's right? (y/n)"
read -r right
if [ -z "$right" ]; then
  echo "Input y or n to continue..."
  exit
fi

if [ "$right" == "n" ]; then
  echo "Fix the wrong paths to continue..."
  exit
fi
echo "----------------------------------------------------------------------"

# Start building
echo "Start building..."
rm -r "${TARGET:?}" && mkdir -p "$TARGET"
cd "$WORKDIR"/cmd/postar || exit

# Go build: windows, linux and darwin
echo "Building windows version..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "$TARGET"/postar-windows.exe

echo "Building linux version..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$TARGET"/postar-linux
chmod +x "$TARGET"/postar-linux

echo "Building darwin version..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "$TARGET"/postar-darwin
chmod +x "$TARGET"/postar-darwin

# Before packaging
echo "Start packaging..."
cd "$TARGET" || exit
cp "$CONFIG_FILE" "$TARGET"/
cp "$LICENSE_FILE" "$TARGET"/

# Start Packaging
echo "Packaging windows version..."
tar -czf postar-$VERSION-windows.tar.gz "$TARGET"/postar-windows.exe "$CONFIG_FILE" "$LICENSE_FILE"
tar -czf postar-$VERSION-linux.tar.gz "$TARGET"/postar-linux "$CONFIG_FILE" "$LICENSE_FILE"
tar -czf postar-$VERSION-darwin.tar.gz "$TARGET"/postar-darwin "$CONFIG_FILE" "$LICENSE_FILE"

# Done
echo "Done."
rm "$TARGET"/postar-windows.exe "$TARGET"/postar-linux "$TARGET"/postar-darwin "$CONFIG_FILE" "$LICENSE_FILE"
