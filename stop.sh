#!/usr/bin/env bash

file="bin/botpid"
if [ ! -e "$file" ]; then
    echo "not running"
    exit 0
fi

kill -9 $(cat "$file")
rm $file
