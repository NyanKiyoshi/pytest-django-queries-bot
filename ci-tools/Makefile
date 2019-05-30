build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/common common.go

clean:
	rm -rf ./bin

run: clean build
	./bin/common
