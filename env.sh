#!/bin/sh

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ -z "$GOPATH" ] ; then
    export GOPATH=$BASE_DIR/go
    export PATH=$PATH:$GOPATH/bin
fi

#Set ENV DEV mode
export ENV=development

#Set DISCORD_TOKEN
export DISCORD_TOKEN=NOP

#Set MONGO_URL
export MONGO_URL="mongodb://mongo:27017"

#Dependencies
go get -d github.com/thommil/fafeliscord
