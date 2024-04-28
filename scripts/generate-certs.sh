#! /usr/bin/env bash

set -e
set -x

pwd

mkdir -p ./dist/certs
mkdir -p ./windows/oem/

[ -f ./dist/certs/tls.key ] || \
openssl req -x509 -newkey ec -pkeyopt ec_paramgen_curve:secp384r1 -days 3650 \
  -nodes -keyout ./dist/certs/tls.key -out ./dist/certs/tls.crt \
  -subj '/C=US/ST=Denial/L=Springfield/O=Dis/CN=anything_but_whitespace' \
  -addext "subjectAltName=DNS:localhost,DNS:*.localhost,DNS:example.org,IP:127.0.0.1,IP:172.17.0.1" \
  -addext 'authorityKeyIdentifier = keyid,issuer'                        \
  -addext 'basicConstraints = CA:FALSE'                                  \
  -addext 'keyUsage = digitalSignature, keyEncipherment'                 \
  -addext 'extendedKeyUsage=serverAuth'

sudo chown 65532:65532 ./dist/certs/*
sudo chmod 444 ./dist/certs/*
sudo rm -rf ./windows/oem/certs
sudo cp -r ./dist/certs ./windows/oem/certs
