.PHONY: build generate-checksum

build:
	mkdir -p dist/
	rm dist/*
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o dist/github github/main.go
	env GOOS=linux go build -ldflags="-s -w" -o dist/uploader uploader/main.go
	env GOOS=linux go build -ldflags="-s -w" -o dist/diff-uploader diff/main.go
	make -C ./ci-tools build

