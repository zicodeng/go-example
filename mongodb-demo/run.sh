#!/usr/bin/env bash

docker rm -f mongodb-demo

docker run -d \
-p 27017:27017 \
--name mongodb-demo \
mongo
