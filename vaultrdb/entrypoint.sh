#!/bin/sh

# ------------------------------------------------------------------------------- setting the variables
# calculating debug mode
if [ -z "$_DEBUG" ]; then
  export _DEBUG="false" 
fi
echo "debug.mode: $_DEBUG" 

# calculate the namespace
if [ ! -f "/var/run/secrets/kubernetes.io/serviceaccount/namespace" ]; then
  export _NAMESPACE="default"
else 
  export _NAMESPACE=$(cat "/var/run/secrets/kubernetes.io/serviceaccount/namespace")
fi

# get the service name
if [ -z "${VAULTRDB_SERVICENAME}" ]; then
  echo "$(hostname)" >  $_TEMPDIR/servicename
  export VAULTRDB_SERVICENAME="$(hostname)"
fi

# ------------------------------------------------------------------------------- service intro
echo "Starting up VaultRDB Service [Version: $(cat '/vaultrdb/VERSION')]"
echo "Operating in Namespace [$_NAMESPACE] as [$VAULTRDB_SERVICENAME]"

# ------------------------------------------------------------------------------- creating the certificate
_TRGTDIR="/tmp/k8s-webhook-server/serving-certs"
_CERTSUBJ="$VAULTRDB_SERVICENAME.$_NAMESPACE.svc"

echo "Creating the Certificate for [$_CERTSUBJ] in [$_TRGTDIR]"

mkdir -p $_TRGTDIR
openssl genrsa -out $_TRGTDIR/tls.key 2048

# print the server.conf into the vaultrdb directory
cat > $_TEMPDIR/server.conf << EOF
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
prompt = no
[req_distinguished_name]
CN = $_CERTSUBJ
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = $_CERTSUBJ
EOF

if [ "$_DEBUG" = "true" ]; then 
  echo "------------------------------------------------------------------------"
  echo "$(cat "$_TEMPDIR/server.conf")"
  echo "------------------------------------------------------------------------"
fi
# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
openssl req -new -key $_TRGTDIR/tls.key -subj "/CN=$_CERTSUBJ" -config $_TEMPDIR/server.conf \
  | openssl x509 -req -CA /vaultrdb/ca.crt -CAkey /vaultrdb/ca.key -CAcreateserial -out $_TRGTDIR/tls.crt -extensions v3_req -extfile $_TEMPDIR/server.conf


# ------------------------------------------------------------------------------- preparing swagger
echo "Preparing the Swagger UI"
_SWAGGERDIR="/vaultrdb/swagger"

# replace the server address fqdn in the $_SWAGGERDIR.yaml for the swaggerui tests
sed -i -e "s/{{FQDN}}/${FQDN}/g" $_SWAGGERDIR/swagger.yaml 
sed -i -e "s/{{VERSION}}/$(cat '/vaultrdb/VERSION')/g" $_SWAGGERDIR/swagger.yaml 
if [ "$_DEBUG" = "true" ]; then 
  echo "------------------------------------------------------------------------"
  echo "$(cat "$_SWAGGERDIR/swagger.yaml")"
  echo "------------------------------------------------------------------------"
fi
# ------------------------------------------------------------------------------- start the main service after sleeping, if neccessary
sleep $SLEEP_BEFORE_SERVICE_START
exec "$@"