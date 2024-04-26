#! /usr/bin/env bash

set -e
set -x

pwd

mkdir -p ./windows/shared/

[ -f ./dockur-sshkey ] || \
ssh-keygen -t rsa -b 4096 -C "dockur" -f dockur-sshkey -q -N ""
mv ./dockur-sshkey.pub ./windows/shared/
sudo chmod 600 ./dockur-sshkey
