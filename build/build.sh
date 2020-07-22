# Postar build script
# Author: fishgoddess
VERSION=v0.1.0-alpha

# Before building
mkdir postar-$VERSION-windows
mkdir postar-$VERSION-linux
mkdir postar-$VERSION-darwin
cd ../src || exit

# Go build: windows, linux and darwin
echo "Building windows version..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../build/postar-$VERSION-windows/postar-$VERSION-windows-amd64.exe main.go
echo "Building linux version..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../build/postar-$VERSION-linux/postar-$VERSION-linux-amd64 main.go
chmod +x ../build/postar-$VERSION-linux/postar-$VERSION-linux-amd64
echo "Building darwin version..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../build/postar-$VERSION-darwin/postar-$VERSION-darwin-amd64 main.go
chmod +x ../build/postar-$VERSION-darwin/postar-$VERSION-darwin-amd64

# Before packaging
cd ../build || exit
mkdir -p ./postar-$VERSION-windows/logs/error
mkdir -p ./postar-$VERSION-linux/logs/error
mkdir -p ./postar-$VERSION-darwin/logs/error
cp ./logit.conf ./postar-$VERSION-windows/
cp ./logit.conf ./postar-$VERSION-linux/
cp ./logit.conf ./postar-$VERSION-darwin/
cp ../_examples/config/postar.english.ini ./postar-$VERSION-windows/postar.ini
cp ../_examples/config/postar.english.ini ./postar-$VERSION-linux/postar.ini
cp ../_examples/config/postar.english.ini ./postar-$VERSION-darwin/postar.ini

# Package all versions
echo "Packaging windows version..."
tar -czf postar-$VERSION-windows-amd64.tar.gz postar-$VERSION-windows
echo "Packaging linux version..."
tar -czf postar-$VERSION-linux-amd64.tar.gz postar-$VERSION-linux
echo "Packaging darwin version..."
tar -czf postar-$VERSION-darwin-amd64.tar.gz postar-$VERSION-darwin

# After packaging
echo "Cleaning..."
rm -rf postar-$VERSION-windows
rm -rf postar-$VERSION-linux
rm -rf postar-$VERSION-darwin
echo "Done."
