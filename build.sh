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
mkdir -p "$TARGET" && rm -rf "${TARGET:?}"/*.tar.gz
cd "$WORKDIR"/cmd/postar || exit

# build builds the target os and arch version package
function build() {
  local GOOS=$1
  local GOARCH=$2
  local BINARY_FILE=$3
  local PKG_FILE="$TARGET"/postar-$VERSION-$GOOS-$GOARCH.tar.gz

  CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$BINARY_FILE"
  tar -czPf "$PKG_FILE" "$BINARY_FILE" "$CONFIG_FILE" "$LICENSE_FILE"
  echo "The $GOOS-$GOARCH package can be found in $PKG_FILE"
  rm "$BINARY_FILE"
}

echo "Building windows-amd64 version..."
build windows amd64 "$TARGET"/postar.exe

echo "Building linux-amd64 version..."
build linux amd64 "$TARGET"/postar

echo "Building darwin-amd64 version..."
build darwin amd64 "$TARGET"/postar

# Done
echo "Done."
