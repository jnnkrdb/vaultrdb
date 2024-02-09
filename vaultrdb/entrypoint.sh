#!/bin/sh

sh /vaultrdb/scripts/environment.sh

echo "Starting up VaultRDB Service [Version: $VAULTRDB_VERSION]"
echo "Operating in Namespace [$VAULTRDB_NAMESPACE] as [$VAULTRDB_SERVICENAME]"
# ------------------------------------------------------------------------------- creating the certificate
sh /vaultrdb/scripts/createCerts.sh
# ------------------------------------------------------------------------------- preparing swagger
sh /vaultrdb/scripts/swagger.sh
# ------------------------------------------------------------------------------- start the main service after sleeping, if neccessary
sleep $SLEEP_BEFORE_SERVICE_START
exec "$@"