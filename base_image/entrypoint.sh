#!/bin/bash

# turn on bash's job control
set -m

gunicorn httpbin:app &

envoy -c envoy-ecds.yaml --log-level info \
    --component-log-level wasm:debug \
    --log-format "[%Y-%m-%d %T.%f][%t][%l][%n] [%g:%#] %v" #--log-path /home/envoy/logs/envoy.log
