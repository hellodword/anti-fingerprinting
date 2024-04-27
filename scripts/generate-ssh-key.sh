#! /usr/bin/env bash

set -e
set -x

pwd

mkdir -p ./windows/shared/

# https://unix.stackexchange.com/questions/69314/automated-ssh-keygen-without-passphrase-how/69318#69318

[ -f ./dockur-sshkey ] || \
ssh-keygen -t rsa -b 4096 -C "dockur" -f dockur-sshkey -q -N ""
mv ./dockur-sshkey.pub ./windows/shared/
sudo chmod 600 ./dockur-sshkey

# # ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i $PWD/dockur-sshkey docker@127.0.0.1 -p 2222 ls
