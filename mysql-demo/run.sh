#!/usr/bin/env bash

docker rm -f mysql-demo

export MYSQL_ROOT_PASSWORD=secret

echo "root password:" $MYSQL_ROOT_PASSWORD

# Run mysql-demo container.
docker run -d \
-p 3306:3306 \
--name mysql-demo \
-e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
-e MYSQL_DATABASE=demo \
zicodeng/mysql-demo
