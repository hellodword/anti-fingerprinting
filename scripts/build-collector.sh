#! /usr/bin/env bash

set -e
set -x

pwd
[ -d .3rd/quic-go ]

go mod download
go build -trimpath -ldflags "-s -w -buildid=" -o dist/collector ./cmd/collector
