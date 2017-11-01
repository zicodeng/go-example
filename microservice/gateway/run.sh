#!/usr/bin/env bash

export ADDR=localhost:3000
export BYE_ADDRS=localhost:4001,localhost:4002
export HELLO_ADDRS=localhost:5001,localhost:5002

go run main.go