#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./CoralBot.go

if [ -n "$1" ]; then
    docker build -t coralbot:$1 ./
fi