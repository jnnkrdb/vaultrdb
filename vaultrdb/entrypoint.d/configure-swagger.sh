#!/bin/sh

set -e

# configuring swagger with the correct domain addresses, version and basic auth if any 

echo "$(date +"%Y-%m-%d - %H:%M:%S") | preparing the swagger ui"

swagger_dir="$VRDB_DIRECTORY_ROOT/swagger"

echo "$(date +"%Y-%m-%d - %H:%M:%S") | configs:"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | -- swagger directory: $swagger_dir"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | ----------- base url: $VRDB_BASE_URL"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | ------------ version: $VRDB_VERSION"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | ----- basicauth user: $BASICAUTH_USER"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | ----- basicauth pass: $BASICAUTH_PASS"

# replace the placeholders in the swagger.yaml files
if [ -d "$VRDB_DIRECTORY_ROOT/swagger/apidocs" ]; then

  find $VRDB_DIRECTORY_ROOT/swagger/apidocs -maxdepth 1 -iname "*.yaml" -type f \
    -exec /bin/sh -c "echo '########################### - {}'" \; \
    -exec /bin/sh -c "sed -i \"s|{{VERSION}}|$VRDB_VERSION|g\" {}" \;
fi

# replace the server address fqdn in the $swagger_dir/swagger.yaml for the swaggerui tests
#sed -i "s|{{BASE_URL}}|$VRDB_BASE_URL|g" $swagger_dir/_swagger.yaml 
#sed -i -e "s|{{BASICAUTH_USER}}|$BASICAUTH_USER|g" $swagger_dir/swagger-initializer.js 
#sed -i -e "s|{{BASICAUTH_PASS}}|$BASICAUTH_PASS|g" $swagger_dir/swagger-initializer.js
 
echo "$(date +"%Y-%m-%d - %H:%M:%S") | swagger ui prepared"