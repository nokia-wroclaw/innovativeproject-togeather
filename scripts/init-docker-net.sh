#!/bin/sh

# Creates Docker bridge network that's used for connect redis and postgres
# to the app's backend

docker network create eathernet || true
