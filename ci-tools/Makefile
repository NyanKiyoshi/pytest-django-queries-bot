build-linux:
	env GOOS=linux go build -ldflags="-s -w" -o bin/queries-upload-linux  ci-upload.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/queries-diff-linux    ci-diff.go

build-macos:
	env GOOS=darwin go build -ldflags="-s -w" -o bin/queries-upload-macos  ci-upload.go
	env GOOS=darwin go build -ldflags="-s -w" -o bin/queries-diff-macos    ci-diff.go

build: build-linux build-macos

clean:
	rm -rf ./bin

run: clean build
	./bin/common
