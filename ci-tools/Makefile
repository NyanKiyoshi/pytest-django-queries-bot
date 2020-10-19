build-linux:
	env GOOS=linux go build -ldflags="-s -w" -o bin/action-handler github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/cmd/action

build-macos:
	env GOOS=darwin go build -ldflags="-s -w" -o bin/action-handler github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/cmd/action

build: clean build-linux build-macos

clean:
	rm -rf ./bin

run: clean build
	./bin/common
