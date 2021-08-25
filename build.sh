# Copyright 2021 Ye Zi Jie.  All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.
#
# Postar build script
# Author: fishgoddess
VERSION=v0.2.1-alpha
BUILD_TARGET=target
CONFIG_FILE=_examples/config/postar.ini

# Before building
echo "Start building to: $BUILD_TARGET"
mkdir $BUILD_TARGET
rm -r ${BUILD_TARGET:?}/bin ${BUILD_TARGET:?}/conf ${BUILD_TARGET:?}/log

# Go build: windows, linux and darwin
echo "Building windows version..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $BUILD_TARGET/bin/postar-$VERSION-windows.exe

echo "Building linux version..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $BUILD_TARGET/bin/postar-$VERSION-linux
chmod +x $BUILD_TARGET/bin/postar-$VERSION-linux

echo "Building darwin version..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $BUILD_TARGET/bin/postar-$VERSION-darwin
chmod +x $BUILD_TARGET/bin/postar-$VERSION-darwin

# Before packaging
echo "Before packaging..."
mkdir -p $BUILD_TARGET/conf
mkdir -p $BUILD_TARGET/log
cp $CONFIG_FILE $BUILD_TARGET/conf/

# Packaging to one
echo "Packaging to one: postar-$VERSION.tar.gz"
cd $BUILD_TARGET || exit
tar -czf postar-$VERSION.tar.gz bin conf log

# Done
echo "Done."
