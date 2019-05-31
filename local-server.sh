#!/usr/bin/env bash

. ./sourceenv.sh

./compile.sh $* || exit $?

# ensure you have installed aws-sam-cli through pip
sam local start-api
