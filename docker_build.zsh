#!/bin/bash
set -e

docker run -itd --rm --name go19 -m 2048M -v $PWD:/app golang:alpine3.16

# docker run -itd --rm --name xo -m 1024M -v $PWD:/app -w /app alpine:3.16
# docker run -itd --rm --name go19 -m 2048M -v $PWD:/app golang:alpine3.16
# ----- proxy & dependency

# https://github.com/mattn/go-sqlite3
# apk add --update gcc musl-dev

# go env -w GO111MODULE=on
# go env -w GOPROXY=https://goproxy.cn,direct