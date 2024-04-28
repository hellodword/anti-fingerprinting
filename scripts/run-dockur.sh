#! /usr/bin/env bash

set -e
set -x

pwd

[ -n "$BROWSER_BROWSER" ]
[ -n "$BROWSER_VERSION" ]

[ -f ./windows/nginx.conf ]

[ "$WINDOWS_VERSION" = "win11" ] || [ "$WINDOWS_VERSION" = "win10" ]

echo "BROWSER_VERSION=$BROWSER_VERSION" >> ./windows/oem/.env
cp ./windows/run-${BROWSER_BROWSER}.ps1 ./windows/oem/run-custom.ps1
cp ./windows/install.bat ./windows/oem/install.bat

docker kill windows || true
docker rm windows || true

docker run --rm -d \
  --stop-timeout 120 --name windows \
  -e MANUAL=N \
  -e VERSION=${WINDOWS_VERSION} \
  -v $PWD/windows/oem:/storage/oem:rw \
  -v $PWD/windows/nginx.conf:/etc/nginx/sites-enabled/web.conf:ro \
  -p 127.0.0.1:2222:22 -p 127.0.0.1:3389:3389 -p 127.0.0.1:8006:8006 \
  --device=/dev/kvm --cap-add NET_ADMIN dockurr/windows:2.22
