.PHONY: build clean deploy

build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/github github/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/uploader uploader/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/diff-uploader diff/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
