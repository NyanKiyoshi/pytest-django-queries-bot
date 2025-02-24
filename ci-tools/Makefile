build-linux:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ../dist/action-handler_amd64_linux github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/cmd/action

build-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ../dist/action-handler_arm64_linux github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/cmd/action

build-macos:
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ../dist/action-handler_amd64_darwin github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/cmd/action

build: build-linux build-linux-arm64 build-macos
