#!/bin/sh

set -e

echo "$(date +"%Y-%m-%d - %H:%M:%S") | starting the entrypoint.sh"

# Execute Startup Scripts
if [ -d "/opt/vaultrdb/entrypoint.d" ]; then

  echo "$(date +"%Y-%m-%d - %H:%M:%S") | executing scripts from [/opt/vaultrdb/entrypoint.d]"

  find /opt/vaultrdb/entrypoint.d -maxdepth 1 -iname "*.sh" -type f \
    -exec /bin/sh -c "echo '########################### - {}'" \; \
    -exec /bin/sh -c "{}" \;
fi

echo "###########################"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | finished entrypoint, starting vaulrdb-bin"

exec $@ 