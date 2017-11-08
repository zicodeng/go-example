#!/usr/bin/env bash

set -e

export BAR_CONTAINER=vh-bar
export VH_NETWORK=vhnet

# Make sure to get the latest image.
docker pull zicodeng/$BAR_CONTAINER

if [ "$(docker ps -aq --filter name=$BAR_CONTAINER)" ]; then
    docker rm -f $BAR_CONTAINER
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

# Run bar Docker container inside our vhnet private network.
docker run \
-d \
--name $BAR_CONTAINER \
--network $VH_NETWORK \
--restart unless-stopped \
zicodeng/$BAR_CONTAINER