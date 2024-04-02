#!/bin/bash

function prepare() {
    mkdir -p "$TARGET" || exit
    mkdir -p "$TARGET"/"$PACKAGE" || exit
}

function build_postar() {
    local binary_file="postar"
    if [[ $GOOS = "windows" ]]; then
        binary_file="$binary_file".exe
    fi

    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$TARGET_PACKAGE"/"$binary_file" "$WORKDIR"/cmd/postar/main.go || exit
    echo "$binary_file"
}

function build_postar_admin() {
    local binary_file="postar-admin"
    if [[ $GOOS = "windows" ]]; then
        binary_file="$binary_file".exe
    fi

    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$TARGET_PACKAGE"/"$binary_file" "$WORKDIR"/cmd/postar-admin/main.go || exit
    echo "$binary_file"
}

function package() {
    cp "$WORKDIR"/config/postar.toml "$TARGET_PACKAGE"/
    cp "$WORKDIR"/config/postar_admin.toml "$TARGET_PACKAGE"/
    cp "$WORKDIR"/LICENSE "$TARGET_PACKAGE"/

    cd "$TARGET"

    local pkg_file="postar-$VERSION-$GOOS-$GOARCH"
    if [[ $GOOS = "windows" ]]; then
        pkg_file="$pkg_file".zip
        zip -qr "$TARGET"/"$pkg_file" "$PACKAGE" || exit
    else
        pkg_file="$pkg_file".tar.gz
        tar -czf "$TARGET"/"$pkg_file" -P "$PACKAGE" || exit
    fi

    cd "$WORKDIR"
    echo "$pkg_file"
}

function clean() {
    rm -rf "$TARGET_PACKAGE"
}

# Main
VERSION=$1
GOOS=$2
GOARCH=$3

WORKDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
TARGET="$WORKDIR"/target
PACKAGE="postar-$VERSION-$GOOS-$GOARCH"
TARGET_PACKAGE="$TARGET"/"$PACKAGE"

echo "-----------------------------------------------------------------------"
echo "VERSION: $VERSION, GOOS: $GOOS, GOARCH:$GOARCH"
echo "WORKDIR: $WORKDIR"
echo "TARGET: $TARGET"
echo "TARGET_PACKAGE: $TARGET_PACKAGE"
echo "-----------------------------------------------------------------------"

# Prepare
prepare || exit
echo "Prepare successfully!"

# Build
postar_binary_file=$(build_postar) || exit
echo "Build $postar_binary_file successfully!"

postar_admin_binary_file=$(build_postar_admin) || exit
echo "Build $postar_admin_binary_file successfully!"

# Package
pkg_file=$(package $postar_binary_file $postar_admin_binary_file) || exit
echo "Package $pkg_file successfully!"

# Done
clean || exit
echo "Done!"
echo ""
