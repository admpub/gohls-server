#!/bin/bash

cp internal/buildinfo/buildinfo.go.in internal/buildinfo/buildinfo.go
go generate github.com/admpub/gohls-server/internal/api
go run *.go ${@:1}