#!/bin/bash

function prepare() {
    mkdir -p "$TARGET" || exit
}

function build_postar() {
    local binary_file="postar"
    if [[ $GOOS = "windows" ]]; then
        binary_file="$binary_file".exe
    fi

    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$TARGET"/"$binary_file" "$WORKDIR"/cmd/postar/main.go || exit
    echo "$binary_file"
}

function build_postar_admin() {
    local binary_file="postar-admin"
    if [[ $GOOS = "windows" ]]; then
        binary_file="$binary_file".exe
    fi

    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$TARGET"/"$binary_file" "$WORKDIR"/cmd/postar-admin/main.go || exit
    echo "$binary_file"
}

function package() {
    local postar_binary_file=$1
    local postar_config_file=postar.toml
    local postar_admin_binary_file=$2
    local postar_admin_config_file=postar_admin.toml
    local license_file=LICENSE

    local pkg_file=postar-$VERSION-$GOOS-$GOARCH.tar.gz
    tar -czf "$TARGET"/"$pkg_file" -C "$TARGET" "$postar_binary_file" -C "$WORKDIR"/configs "$postar_config_file" -C "$TARGET" "$postar_admin_binary_file" -C "$WORKDIR"/configs "$postar_admin_config_file" -C "$WORKDIR" "$license_file" || exit

    echo "$pkg_file"
}

function clean() {
    local postar_binary_file=$1
    local postar_admin_binary_file=$2

    rm "$TARGET"/"$postar_binary_file"
    rm "$TARGET"/"$postar_admin_binary_file"
}

VERSION=$1
GOOS=$2
GOARCH=$3

WORKDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
TARGET=$WORKDIR/target

echo "-----------------------------------------------------------------------"
echo "VERSION: $VERSION, GOOS: $GOOS, GOARCH:$GOARCH"
echo "WORKDIR: $WORKDIR"
echo "TARGET: $TARGET"
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
clean $postar_binary_file $postar_admin_binary_file || exit
echo "Done!"
