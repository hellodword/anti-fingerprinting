#! /usr/bin/env bash

set -e
set -x

pwd
[ -d .3rd/quic-go ] || \
  git clone --depth 1 -b v0.42.0 https://github.com/quic-go/quic-go .3rd/quic-go
pushd .3rd/quic-go
git apply ../../quic-go.patch
popd
