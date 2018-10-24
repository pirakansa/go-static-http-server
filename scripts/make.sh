#!/bin/bash
# Copyright (c) 2018 pirakansa

CMD_DIR=$(cd $(dirname $0);cd ../cmd/static-http-server/;pwd)
GOBUILD="go build"
GORUN="go run"
GOASSETSBUILD="go-assets-builder"
GOCLEAN="go clean"
GOTEST="go test"
GOGET="dep ensure"

function __Build() {
    cd ${CMD_DIR}
    ${GOASSETSBUILD} -p spage -o ../../spage/assets.go -s="/html" html/
    ${GOBUILD}
}

function __Debug() {
    cd ${CMD_DIR}
    ${GORUN} -tags=debug main.go
}

function __Clean() {
    cd ${CMD_DIR}
    ${GOCLEAN}
}

function __Install() {
    cd ${CMD_DIR}
    ${GOGET}
}

case "$1" in
    "clean" ) 
        __Clean ;;
    "install" ) 
        __Install ;;
    "debug" ) 
        __Debug ;;
    * ) 
        __Build ;;
esac

