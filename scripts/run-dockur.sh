#! /usr/bin/env bash

set -e
set -x

pwd

[ -n "$BROWSER_BROWSER" ]
[ -n "$BROWSER_VERSION" ]

[ -f ./windows/nginx.conf ]

[ "$WINDOWS_VERSION" = "win11" ] || [ "$WINDOWS_VERSION" = "win10" ]

echo "BROWSER_VERSION=$BROWSER_VERSION" >> ./windows/shared/.env
cp ./windows/run-${BROWSER_BROWSER}.ps1 ./windows/shared/run-custom.ps1

docker kill windows || true
docker rm windows || true

docker run --rm -d \
  --stop-timeout 120 --name windows \
  -e MANUAL=N \
  -e VERSION=${WINDOWS_VERSION} \
  -v $PWD/windows/shared:/storage/shared:rw \
  -v $PWD/windows/${WINDOWS_VERSION}x64.xml:/run/assets/${WINDOWS_VERSION}x64.xml:ro \
  -v $PWD/windows/nginx.conf:/etc/nginx/sites-enabled/web.conf:ro \
  -p 127.0.0.1:2222:22 -p 127.0.0.1:3389:3389 -p 127.0.0.1:8006:8006 \
  --device=/dev/kvm --cap-add NET_ADMIN dockurr/windows
