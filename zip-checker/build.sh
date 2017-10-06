#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t zicodeng/zip-checker .
docker run -d -e ADDR=:3000 -p 3000:3000 --name zip-checker zicodeng/zip-checker