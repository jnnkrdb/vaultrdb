#!/bin/sh

set -e

# creating the required certs
# for the webhook and the
# other service containers
# connected to this operator

echo "$(date +"%Y-%m-%d - %H:%M:%S") | generating required certs for k8s webhook"

targetDir="/tmp/k8s-webhook-server/serving-certs"
certSubject="$VRDB_SERVICE_NAME.$VRDB_SERVICE_NAMESPACE.svc"

mkdir -p $targetDir

openssl genrsa -out $targetDir/tls.key 2048

# exchange the placeholders from 
# the server.conf with the actual 
# values
sed -i "s/{{CERTIFICATESUBJECT}}/$certSubject/g" $VRDB_DIRECTORY_ROOT/config/certs/server.conf

openssl req \
  -new \
  -key $targetDir/tls.key \
  -subj "/CN=$certSubject" \
  -config $VRDB_DIRECTORY_ROOT/config/certs/server.conf | openssl x509 \
    -req \
    -CA $VRDB_DIRECTORY_ROOT/config/certs/ca.crt \
    -CAkey $VRDB_DIRECTORY_ROOT/config/certs/ca.key \
    -CAcreateserial \
    -out $targetDir/tls.crt \
    -extensions v3_req \
    -extfile $VRDB_DIRECTORY_ROOT/config/certs/server.conf

