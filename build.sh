#!/bin/sh

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ -z "$GOPATH" ] ; then
    export GOPATH=$BASE_DIR/go
    export PATH=$PATH:$GOPATH/bin
fi    

#build
go build -i -o $GOPATH/bin/fafeliscord $BASE_DIR/main.go