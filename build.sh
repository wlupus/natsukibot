#!/usr/bin/env bash
go build -o natsukibot src/main.go

if [ ! -d bin ] ; then
    mkdir bin
fi

mv natsukibot bin/

