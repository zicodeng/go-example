#!/usr/bin/env bash
set -e

# Build gateway Docker container.
./build.sh

export APP_NETWORK=vhnet
export FOO_ADDR=vh-foo:80
export BAR_ADDR=vh-bar:80

if ! [ "$(docker network ls | grep $APP_NETWORK)" ]; then
    docker network create $APP_NETWORK
fi

docker run \
-d \
-p 80:80 \
-e FOO_ADDR=$FOO_ADDR \
-e BAR_ADDR=$BAR_ADDR \
--network $APP_NETWORK \
--name vh-gateway \
zicodeng/vh-gateway
