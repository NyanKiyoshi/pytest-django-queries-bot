## Dependencies
1. npm
1. go (`brew install go`)

## Install (first run)
1. `npm i`
1. `export PATH="$PATH:$(npm bin)"`

## Build
1. `make build`

## Local Testing
1. install docker
1. pip install aws-sam-cli
1. `sam local start-api`

## Deploy
1. `sls deploy`
