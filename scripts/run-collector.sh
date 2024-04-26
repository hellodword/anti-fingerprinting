#! /usr/bin/env bash

set -e
set -x

pwd

mkdir -p ./dist

[ -f ./dist/certs/tls.key ]
[ -f ./dist/collector ]

mkdir -p ./dist/db
[ "$(stat -c '%u' ./dist/db)" == "65532" ] || sudo chown 65532:65532 ./dist/db
[ "$(stat -c '%g' ./dist/db)" == "65532" ] || sudo chown 65532:65532 ./dist/db

cp ./cmd/collector/docker-compose.yml ./dist/

cd dist
docker compose down
docker compose up --build -d
