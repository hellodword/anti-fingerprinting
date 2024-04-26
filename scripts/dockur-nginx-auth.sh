#! /usr/bin/env bash

set -e
set -x

pwd

mkdir -p ./windows/shared/

echo -n 123 | openssl passwd -apr1 -stdin 

if [ -n "$NGINX_AUTH" ]; then
  echo "admin:$(echo -n "$NGINX_AUTH" | openssl passwd -apr1 -stdin)" > ./windows/shared/.htpasswd
else
  echo 'admin:$apr1$zyjICK.T$K6Q9fLUxiVMN2VQcUlDAI.' > ./windows/shared/.htpasswd
fi


curl -fsSL --output ./windows/nginx.conf https://github.com/qemus/qemu-docker/raw/master/web/nginx.conf
sed -i '2s@^@    auth_basic "Administrator’s Area";\n    auth_basic_user_file /storage/shared/.htpasswd;\n@' ./windows/nginx.conf
head ./windows/nginx.conf
