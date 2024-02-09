#!/bin/sh

echo "Creating the Webhook Certificates from the ca.crt"

# setting the variables
_TEMPDIR="/vaultrdb/temp"
_SWAGGERDIR="/vaultrdb/swagger"
_TRGTDIR="/tmp/k8s-webhook-server/serving-certs"
_CERTSUBJ="$VAULTRDB_SERVICENAME.$VAULTRDB_NAMESPACE.svc"

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

echo "------------------------------------------------------------------------"
echo "$(cat "$_TEMPDIR/server.conf")"
echo "------------------------------------------------------------------------"

# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
openssl req -new -key $_TRGTDIR/tls.key -subj "/CN=$_CERTSUBJ" -config $_TEMPDIR/server.conf \
  | openssl x509 -req -CA /vaultrdb/ca.crt -CAkey /vaultrdb/ca.key -CAcreateserial -out $_TRGTDIR/tls.crt -extensions v3_req -extfile $_TEMPDIR/server.conf
