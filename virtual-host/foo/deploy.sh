#!/usr/bin/env bash
set -e

./build.sh

docker push zicodeng/vh-foo

ssh -oStrictHostKeyChecking=no root@104.236.160.50 'bash -s' < run.sh