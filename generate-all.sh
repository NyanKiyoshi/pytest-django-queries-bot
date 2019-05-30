#!/usr/bin/env bash

HERE=$(dirname $0)
cd $HERE

DIRS="uploader"
IFS=" "

for dir in $DIRS; do (
        echo -n "-> generating... $dir/go.gen ..."
        { cd $dir && go run gen.go ; } || { echo "Failed..." >&2 ; exit 1 ; }
        echo "ok"
    )
done
