.PHONY:clean
VERSION=v0.3.1-alpha
CONFIG_FILE=./_examples/config/postar.ini
test:
	go test -v -cover ./...
run:
	go run cmd/postar/main.go -config.file $(CONFIG_FILE)
linux:
	./build.sh linux amd64 postar $(VERSION)
windows:
	./build.sh windows amd64 postar.exe $(VERSION)
darwin:
	./build.sh darwin amd64 postar $(VERSION)
build:
	make linux && make windows && make darwin
clean:
	rm -rf ./target
all:
	make test && make clean && make build