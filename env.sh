#!/bin/sh

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ -z "$GOPATH" ] ; then
    export GOPATH=$BASE_DIR/go
    export PATH=$PATH:$GOPATH/bin
fi

#Set ENV DEV mode
export ENV=development

#Set DISCORD_TOKEN
export DISCORD_TOKEN=NjQyMzQ2NjYxNzE2Mjk1Njkx.XcVmZA.g1KlRbu98Lp_XjlBPnFPoCDQV6A

#Dependencies
go get -d github.com/thommil/fafeliscord
