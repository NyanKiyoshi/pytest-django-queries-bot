build-linux:
	env GOOS=linux go build -ldflags="-s -w" -o bin/queries-upload             ci-upload.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/queries-diff               ci-diff.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/queries-get-gh-base-ref    ci-github-base-ref.go

build-macos:
	env GOOS=darwin go build -ldflags="-s -w" -o bin/queries-upload-macos            ci-upload.go
	env GOOS=darwin go build -ldflags="-s -w" -o bin/queries-diff-macos              ci-diff.go
	env GOOS=darwin go build -ldflags="-s -w" -o bin/queries-get-gh-base-ref-macos   ci-github-base-ref.go

build: clean build-linux build-macos

clean:
	rm -rf ./bin

run: clean build
	./bin/common
