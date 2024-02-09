#!/bin/sh

echo "Preparing the Swagger UI"

# setting the variables
_TEMPDIR="/vaultrdb/temp"
_SWAGGERDIR="/vaultrdb/swagger"

# replace the server address fqdn in the $_SWAGGERDIR.yaml for the swaggerui tests
sed -i -e "s/{{FQDN}}/${FQDN}/g" $_SWAGGERDIR/swagger.yaml > $_TEMPDIR/swagger.temp
cp $_TEMPDIR/swagger.temp $_SWAGGERDIR/swagger.yaml
echo "------------------------------------------------------------------------"
echo "$(cat "$_SWAGGERDIR/swagger.yaml")"
echo "------------------------------------------------------------------------"
