#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

##
## Runs Envoy and the example control plane server.  See
## `internal/example` for the go source.
##

# Envoy start-up command  - it suppose envoy path in our system is '/usr/local/bin/envoy' we can copy this file from envoy docker to our local system bin
ENVOY=${ENVOY:-/usr/local/bin/envoy}

# Start envoy: important to keep drain time short - connect to our grpc server on through envoy configuration
(${ENVOY} -c sample/bootstrap-xdsv3.yaml --drain-time-s 1 -l debug)&
ENVOY_PID=$!

function cleanup() {
  kill ${ENVOY_PID}
}
trap cleanup EXIT

# Run the control plane - run a grpc server on this ip 127.0.0.1:18000
bin/example -debug $@  #$@: The filename representing the target. #https://stackoverflow.com/a/37701195/581476
