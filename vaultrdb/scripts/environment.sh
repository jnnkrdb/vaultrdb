#!/bin/sh

# set the version of vaultrdb
export VAULTRDB_VERSION=$(cat '/vaultrdb/VERSION')

# calculate the namespace
if [ -f "/var/run/secrets/kubernetes.io/serviceaccount/namespace" ]; then
  export VAULTRDB_NAMESPACE=$(cat "/var/run/secrets/kubernetes.io/serviceaccount/namespace")
else
  export VAULTRDB_NAMESPACE="default"
fi

# get the service name
if [ -z "${VAULTRDB_SERVICENAME}" ]; then
  export VAULTRDB_SERVICENAME=$(hostname)
fi