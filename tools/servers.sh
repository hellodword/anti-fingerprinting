#! /usr/bin/env bash

set -e
set -x

# # inotifywait -e modify result.json

mkdir -p certs
[ -f certs/tls.key ] || \
openssl req -x509 -newkey ec -pkeyopt ec_paramgen_curve:secp384r1 -days 3650 \
  -nodes -keyout certs/tls.key -out certs/tls.crt -subj "/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,DNS:*.localhost,IP:127.0.0.1"

if [ ! -d tmp/TrackMe ]; then
  git clone --depth 1 https://github.com/wwhtrbbtt/TrackMe tmp/TrackMe
  git -C tmp/TrackMe apply ../../TrackMe.patch
fi

pushd tmp/TrackMe
docker build -t trackme .
popd

[ -f trackme.result.json ] || touch trackme.result.json

# # https://127.0.0.1:8443/api/all
# docker run --rm -d \
#   --name trackme \
#   -p 8443:443 \
#   -v `pwd`/TrackMe.json:/app/config.json:ro \
#   -v `pwd`/certs:/app/certs:ro \
#   -v `pwd`/trackme.result.json:/app/result.json:rw \
#   trackme


if [ ! -d tmp/fingerproxy ]; then
  git clone --depth 1 https://github.com/wi1dcard/fingerproxy tmp/fingerproxy
  git -C tmp/fingerproxy apply ../../fingerproxy.patch
fi

pushd tmp/fingerproxy
docker build -t fingerproxy .
popd

[ -f fingerproxy.result.json ] || touch fingerproxy.result.json

# # https://127.0.0.1:9443/json
# docker run --rm -d \
#   --name fingerproxy \
#   -p 9443:9443 \
#   -v `pwd`/certs:/app/certs:ro \
#   -v `pwd`/fingerproxy.result.json:/app/result.json:rw \
#   fingerproxy

