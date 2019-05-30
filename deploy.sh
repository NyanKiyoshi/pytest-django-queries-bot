#!/usr/bin/env bash

. ./sourceenv.sh

# hint: use ./deploy.sh --force if you want to force deploy
sls deploy $*
