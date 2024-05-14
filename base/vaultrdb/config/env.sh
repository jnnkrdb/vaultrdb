#!/bin/sh

echo "$(date +"%Y-%m-%d - %H:%M:%S") | setting the environment variables"

# --------------------------------------------------------------------------- directory sets
export VRDB_DIRECTORY_ROOT        ="/opt/vaultrdb"

# --------------------------------------------------------------------------- dns configs
if [ ! -f "/var/run/secrets/kubernetes.io/serviceaccount/namespace" ]; then
  export VRDB_SERVICE_NAMESPACE="default"
else 
  export VRDB_SERVICE_NAMESPACE=$(cat "/var/run/secrets/kubernetes.io/serviceaccount/namespace")
fi

# get the service name
if [ -z "${VRDB_SERVICE_NAME}" ]; then
  export VRDB_SERVICE_NAME="$(hostname)"
fi

# --------------------------------------------------------------------------- general configs
export VRDB_VERSION="$(cat "$VRDB_DIRECTORY_ROOT/config/VERSION")"

# --------------------------------------------------------------------------- web configs
if [ -z "${VRDB_BASE_URL}" ]; then
  export VRDB_BASE_URL="http://localhost:80/"
fi

if [ -z "${VRDB_BASICAUTH_USER}" ]; then
  export VRDB_BASICAUTH_USER="vault"
fi

if [ -z "${VRDB_BASICAUTH_PASSWORD}" ]; then
  export VRDB_BASICAUTH_PASSWORD="vault"
fi

if [ -z "${VRDB_ENABLE_SWAGGER}" ]; then
  export VRDB_ENABLE_SWAGGER="false"
fi

echo "$(date +"%Y-%m-%d - %H:%M:%S") | environment variables set"