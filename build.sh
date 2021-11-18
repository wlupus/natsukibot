#!/usr/bin/env bash
go build -o wlupusbot src/main.go

if [ ! -d bin ] ; then
    mkdir bin
fi

mv wlupusbot bin/

