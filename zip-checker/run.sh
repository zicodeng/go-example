#!/usr/bin/env bash
docker rm -f zip-checker

docker run -d \
-p 443:443 \
--name zip-checker \
-v /c/Users/Zico\ Deng/Desktop/go/src/github.com/zicodeng/go-example/zip-checker/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
zicodeng/zip-checker