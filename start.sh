#!/usr/bin/env bash

if [ ! -e "bin/natsukibot" ]; then
    echo "run ./build.sh first"
    exit 1
fi

./bin/natsukibot &> /dev/null &
echo "$!" > bin/botpid
