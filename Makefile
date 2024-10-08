.PHONY: all fmt test build clean proto postar postaradmin

VERSION=v0.5.1

all:
	make test && make clean && make build

fmt:
	go fmt ./...

test:
	go mod tidy
	go test -cover ./...

build:
	go mod tidy
	./build.sh $(VERSION) linux amd64
	./build.sh $(VERSION) darwin amd64
	./build.sh $(VERSION) windows amd64

clean:
	rm -rf ./target

proto:
	cd api && buf build && buf generate

postar:
	go mod tidy
	go run cmd/postar/main.go -conf ./config/postar.toml

postaradmin:
	go mod tidy
	go run cmd/postar-admin/main.go -conf ./config/postar_admin.toml
