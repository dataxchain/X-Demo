#!/bin/sh
set -e

ipfs init

# config the api endpoint, may introduce security risk to expose API_PORT public
ipfs config Addresses.API /ip4/0.0.0.0/tcp/${API_PORT}
# config the gateway endpoint
ipfs config Addresses.Gateway /ip4/0.0.0.0/tcp/${GATEWAY_PORT}

ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["*"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT", "GET", "POST"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Credentials '["true"]'

exec "$@"
