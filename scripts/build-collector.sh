#! /usr/bin/env bash

set -e
set -x

pwd
[ -d .3rd/quic-go ]

go mod download
go build -trimpath -tags "patched" -buildvcs=false -ldflags "-s -w -buildid=" -o dist/collector ./cmd/collector
