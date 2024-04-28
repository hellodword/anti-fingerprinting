#! /usr/bin/env bash

set -e
set -x

pwd

mkdir -p ./windows/oem/

echo -n 123 | openssl passwd -apr1 -stdin 

if [ -n "$NGINX_AUTH" ]; then
  echo "admin:$(echo -n "$NGINX_AUTH" | openssl passwd -apr1 -stdin)" > ./windows/oem/.htpasswd
else
  echo 'admin:$apr1$zyjICK.T$K6Q9fLUxiVMN2VQcUlDAI.' > ./windows/oem/.htpasswd
fi

# Authentication: https://github.com/dockur/windows/issues/301#issuecomment-2018610554
curl -fsSL --output ./windows/nginx.conf https://github.com/qemus/qemu-docker/raw/master/web/nginx.conf
sed -i '2s@^@    auth_basic "Administrator’s Area";\n    auth_basic_user_file /storage/oem/.htpasswd;\n@' ./windows/nginx.conf
head ./windows/nginx.conf
