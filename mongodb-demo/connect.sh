#!/usr/bin/env bash

docker run -it \
--rm \
--network host \
mongo sh -c 'exec mongo 127.0.0.1/demo'
