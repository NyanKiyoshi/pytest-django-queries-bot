build-linux:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/action-handler_amd64_linux github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/cmd/action

build-macos:
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/action-handler_amd64_darwin github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/cmd/action

build: clean build-linux build-macos

clean:
	rm -rf ./bin

run: clean build
	./bin/common
