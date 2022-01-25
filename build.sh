#!/bin/bash

VERSION=v0.3.0-alpha
echo "VERSION: $VERSION"
echo "----------------------------------------------------------------------"

# Check
WORKDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
echo "WORKDIR: $WORKDIR"
CONFIG_DIR=$WORKDIR/_examples/config
CONFIG_FILE=postar.ini
echo "CONFIG: $CONFIG_DIR/$CONFIG_FILE"
LICENSE_FILE=LICENSE
echo "LICENSE: $WORKDIR/$LICENSE"
TARGET=$WORKDIR/target
echo "TARGET: $TARGET"
echo "----------------------------------------------------------------------"

echo "Want to continue? (y/n)"
read -r right
if [ -z "$right" ]; then
  echo "Input y or n to continue..."
  exit
fi

if [ "$right" == "n" ]; then
  echo "Fix the problems to continue..."
  exit
fi
echo "----------------------------------------------------------------------"

# Prepare
echo "Preparing..."
mkdir -p "$TARGET" && rm -rf "${TARGET:?}"/*.tar.gz || exit
cd "$WORKDIR"/cmd/postar || exit

# build builds the target os and arch version package
function build() {
  local GOOS=$1
  local GOARCH=$2
  local BINARY_FILE=$3
  local PKG_FILE="$TARGET"/postar-$VERSION-$GOOS-$GOARCH.tar.gz

  CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$TARGET"/"$BINARY_FILE" || exit
  tar -czf "$PKG_FILE" -C "$TARGET" "$BINARY_FILE" -C "$CONFIG_DIR" "$CONFIG_FILE" -C "$WORKDIR" "$LICENSE_FILE" || exit
  echo "The $GOOS-$GOARCH package can be found in $PKG_FILE" || exit
  rm "$TARGET"/"$BINARY_FILE" || exit
}

echo "Building windows-amd64 version..."
build windows amd64 postar.exe

echo "Building linux-amd64 version..."
build linux amd64 postar

echo "Building darwin-amd64 version..."
build darwin amd64 postar

# Done
echo "Done."
