#!/usr/bin/env bash

if [ ! -e "bin/wlupusbot" ]; then
    echo "run ./build.sh first"
    exit 1
fi

./bin/wlupusbot &> /dev/null &
echo "$!" > bin/botpid
