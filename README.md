# VaultRDB

## Install to Cluster

To install the operator, use the manifests from the `.doc/deploy` directory.

### Using own CA Certificate

If you want to use your own CA certificate, you have to create a `ca.crt` and a `ca.key` and mount them into the operator container under `/vaultrdb/ca.crt` and `/vaultrdb/ca.key`.
If you don't have an own certificate, you can create one with the following commands:#
```bash
openssl req -nodes -new -x509 -keyout ca.key -out ca.crt -subj "/CN=Webhook Certification for VaultRDB"
echo "$(openssl base64 -A <"ca.crt")"
```
