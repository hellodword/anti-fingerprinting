#! /usr/bin/env bash

set -e
set -x

pwd

go mod download
go build -trimpath -buildvcs=false -ldflags "-s -w -buildid=" -o dist/compare-collectorinfo ./cmd/compare-collectorinfo
