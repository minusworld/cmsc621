#!/bin/sh

export GOPATH=`pwd`
go run ./src/server/main.go $1
