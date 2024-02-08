#!/bin/sh

echo "Starting up VaultRDB Service [Version: $(cat '/vaultrdb/VERSION')]"


if [ -f "/var/run/secrets/kubernetes.io/serviceaccount/namespace" ]; then
  NAMESPACE=$(cat "/var/run/secrets/kubernetes.io/serviceaccount/namespace")
else
  NAMESPACE="default"
fi

echo "Operating in Namespace [$NAMESPACE]"

# ------------------------------------------------------------------------------- creating the certificate
[[ ! -z "$SERVICENAME" ]] && SERVICENAME=$(hostname)

_CERTSUBJ="$SERVICENAME.$NAMESPACE.svc"
_TRGTDIR="/tmp/k8s-webhook-server/serving-certs"

echo "Creating the Certificate for [$_CERTSUBJ] in [$_TRGTDIR]"

mkdir -p $_TRGTDIR
openssl genrsa -out $_TRGTDIR/tls.key 2048

# print the server.conf into the vaultrdb directory
cat > /vaultrdb/server.conf << EOF
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

echo "------------------------------------------------------------------------"
echo "$(cat '/vaultrdb/server.conf')"
echo "------------------------------------------------------------------------"

# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
openssl req -new -key $_TRGTDIR/tls.key -subj "/CN=$_CERTSUBJ" -config /vaultrdb/server.conf \
  | openssl x509 -req -CA /vaultrdb/ca.crt -CAkey /vaultrdb/ca.key -CAcreateserial -out $_TRGTDIR/tls.crt -extensions v3_req -extfile /vaultrdb/server.conf

# ------------------------------------------------------------------------------- start the main service after sleeping, if neccessary
sleep $SLEEP_BEFORE_SERVICE_START
exec "$@"

