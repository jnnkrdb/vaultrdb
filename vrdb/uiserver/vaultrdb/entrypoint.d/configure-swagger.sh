#!/bin/sh

set -e

# configuring swagger with the correct domain addresses, version and basic auth if any 

echo "$(date +"%Y-%m-%d - %H:%M:%S") | preparing the swagger ui"

swagger_dir="$VRDB_DIRECTORY_ROOT/web/swagger"

# replace the server address fqdn in the $swagger_dir/swagger.yaml for the swaggerui tests
sed -i -e "s/{{BASE_URL}}/$VRDB_BASE_URL/g" $swagger_dir/swagger.yaml 
sed -i -e "s/{{VERSION}}/$VRDB_VERSION/g" $swagger_dir/swagger.yaml 
#sed -i -e "s/{{BASICAUTH_USER}}/${BASICAUTH_USER}/g" $swagger_dir/swagger-initializer.js 
#sed -i -e "s/{{BASICAUTH_PASS}}/${BASICAUTH_PASS}/g" $swagger_dir/swagger-initializer.js
