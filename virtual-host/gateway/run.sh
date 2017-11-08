#!/usr/bin/env bash

set -e

export GATEWAY_CONTAINER=vh-gateway
export FOO_CONTAINER=vh-foo
export BAR_CONTAINER=vh-bar
export VH_NETWORK=vhnet

export FOO_ADDR=$FOO_CONTAINER:80
export BAR_ADDR=$BAR_CONTAINER:80

# Make sure to get the latest image.
docker pull zicodeng/$GATEWAY_CONTAINER

# Remove the old containers first.
if [ "$(docker ps -aq --filter name=$GATEWAY_CONTAINER)" ]; then
    docker rm -f $GATEWAY_CONTAINER
fi


# Remove dangling images.
if [ "$(docker images -q -f dangling=true)" ]; then
    docker rmi $(docker images -q -f dangling=true)
fi

# Clean up the system.
docker system prune -f

# Create Docker private network if not exist.
if ! [ "$(docker network ls | grep $VH_NETWORK)" ]; then
    docker network create $VH_NETWORK
fi

# Run gateway Docker container inside our vhnet private network.
# Our gateway container will be the only container that has port published.
docker run \
-d \
-p 80:80 \
--name $GATEWAY_CONTAINER \
--network $VH_NETWORK \
-e FOO_ADDR=$FOO_CONTAINER \
-e BAR_ADDR=$BAR_CONTAINER \
--restart unless-stopped \
zicodeng/$GATEWAY_CONTAINER