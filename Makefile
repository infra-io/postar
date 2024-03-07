.PHONY: test fmt proto postar postaradmin linux windows darwin build clean all

VERSION=v0.4.0-alpha

test:
	go mod tidy
	go test -cover ./...

fmt:
	go fmt ./...

proto:
	cd api && buf build && buf generate

postar:
	go mod tidy
	go run cmd/postar/main.go -conf ./configs/postar.toml

postaradmin:
	go mod tidy
	go run cmd/postar-admin/main.go -conf ./configs/postar_admin.toml

build:
	go mod tidy
	./build.sh $(VERSION) linux amd64
	./build.sh $(VERSION) darwin amd64
	./build.sh $(VERSION) windows amd64

clean:
	rm -rf ./target

all:
	make test && make clean && make build