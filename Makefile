.PHONY: all build dist test install clean tools deps update-deps

all:
	@echo "build:      Build code."
	@echo "test:       Run tests."
	@echo "install:    Install binary."
	@echo ""
	@echo "clean:      Clean up."

build:
	@mkdir -p ./bin
	@rm -f ./bin/*
	go build -o ./bin/core ./cmd/core

dist:
	@mkdir -p ./bin
	@rm -f ./bin/*
	GOOS=darwin GOARCH=amd64 go build -o ./bin/core-darwin64 ./cmd/core
	GOOS=linux GOARCH=amd64 go build -o ./bin/core-linux64 ./cmd/core
	tar -czf ./bin/core-linux64.tar.gz ./bin/core-linux64
	tar -czf ./bin/core-darwin64.tar.gz ./bin/core-darwin64

test:
	go test ./...

install: build
	go install ./...

clean:
	@rm -rf ./bin