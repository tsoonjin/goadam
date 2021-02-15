#!/usr/bin/env bash
# Bash3 Boilerplate. Copyright (c) 2014, kvz.io

set -o errexit
set -o pipefail
set -o nounset

LOCAL_PORT=8000
APP_PORT=7000
APP_NAME=goadam

docker build --build-arg APP_PORT=$APP_PORT --build-arg APP_NAME=$APP_NAME -t $APP_NAME . &&

docker run -p $LOCAL_PORT:$APP_PORT --env-file .env $APP_NAME
