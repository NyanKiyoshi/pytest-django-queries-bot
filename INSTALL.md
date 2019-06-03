## Dependencies
1. npm
1. go (`brew install go`)
1. go-dep (`brew install dep` or `apt-get install go-dep`)

## Install (first run)
1. `npm i`
1. `export PATH="$PATH:$(npm bin)"`

## Build
1. `./compile`

## Local Testing
1. install docker
1. pip install aws-sam-cli
1. `./local-server.sh`

## Deploy
1. `serverless deploy`
