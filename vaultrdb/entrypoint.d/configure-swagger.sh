#!/bin/sh

set -e

# configuring swagger with the correct domain addresses, version and basic auth if any 

echo "$(date +"%Y-%m-%d - %H:%M:%S") | preparing the swagger ui"

echo "$(date +"%Y-%m-%d - %H:%M:%S") | configs:"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | -- swagger directory: /opt/vaultrdb/swagger"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | ----------- base url: /opt/vaultrdb"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | ------------ version: $(cat /opt/vaultrdb/home/VERSION)"

# replace the placeholders in the swagger.yaml files
if [ -d "/opt/vaultrdb/swagger/apidocs" ]; then
  echo "$(date +"%Y-%m-%d - %H:%M:%S") | preparing:"

  find /opt/vaultrdb/swagger/apidocs -maxdepth 1 -iname "*.yaml" -type f \
    -exec /bin/sh -c "echo '- {}'" \; \
    -exec /bin/sh -c "sed -i \"s|{{VERSION}}|$(cat /opt/vaultrdb/home/VERSION)|g\" {}" \;
fi

echo "$(date +"%Y-%m-%d - %H:%M:%S") | swagger ui prepared"