#! /usr/bin/env bash

set -e
set -x

pwd
[ -d .3rd/quic-go ] || \
  git clone --depth 1 -b v0.44.0 https://github.com/quic-go/quic-go .3rd/quic-go
pushd .3rd/quic-go
[ -f .patched ] || \
(git apply ../../patches/quic-go.patch && touch .patched)
popd
go mod edit -replace=github.com/quic-go/quic-go@v0.44.0=./.3rd/quic-go
go mod tidy
