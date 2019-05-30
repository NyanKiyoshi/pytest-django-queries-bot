#!/usr/bin/env bash

. ./sourceenv.sh

bash ./generate-all.sh || { echo "Failed generating files..." ; exit 1 ; }

make $*
