#!/usr/bin/env bash

docker run -it \
--rm \
--network host \
zicodeng/mysql-demo sh -c "mysql -h127.0.0.1 -uroot -p$MYSQL_ROOT_PASSWORD"