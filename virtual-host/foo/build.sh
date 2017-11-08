#!/usr/bin/env bash

export CONTAINER_NAME=vh-foo

GOOS=linux go build

docker build -t zicodeng/$CONTAINER_NAME .

if [ "$(docker ps -aq --filter name=$CONTAINER_NAME)" ]; then
    docker rm -f $CONTAINER_NAME
fi

# Remove dangling images.
if [ "$(docker images -q -f dangling=true)" ]; then
    docker rmi $(docker images -q -f dangling=true)
fi

go clean