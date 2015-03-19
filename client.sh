#!/bin/sh

export GOPATH=`pwd`
go run ./src/client/main.go $1
