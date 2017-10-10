#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t zicodeng/zip-checker .
docker push zicodeng/zip-checker
go clean