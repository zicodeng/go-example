#!/usr/bin/env bash

./build.sh

docker push zicodeng/vh-gateway

# Send run.sh to the cloud running remotely.
ssh -oStrictHostKeyChecking=no root@104.236.160.50 'bash -s' < run.sh