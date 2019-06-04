.PHONY: build clean deploy

generate:
	env go generate gen.go

build: generate
	dep ensure -v
	env GOOS=darwin go build -ldflags="-s -w" -o bin/ghauthorize ghapp/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/github github/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/uploader uploader/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/diff-uploader diff/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
