#!/bin/sh

export GOPROXY=https://goproxy.cn
dlv --headless --log --listen :9080 --api-version 2 --accept-multiclient debug main.go