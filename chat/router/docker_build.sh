#!/usr/bin/env bash

prog=chatrouter

git pull
docker rm -f $prog
docker rmi -f $prog
set -e
docker build --rm -t $prog .
