#!/bin/sh

set -e

echo "$(date +"%Y-%m-%d - %H:%M:%S") | starting the entrypoint.sh"

. /opt/vaultrdb/config/env.sh

# Execute Startup Scripts
if [ -d "$VRDB_DIRECTORY_ROOT/entrypoint.d" ]; then

  echo "$(date +"%Y-%m-%d - %H:%M:%S") | executing scripts from [$VRDB_DIRECTORY_ROOT/entrypoint.d]"

  find $VRDB_DIRECTORY_ROOT/entrypoint.d -maxdepth 1 -iname "*.sh" -type f \
    -exec /bin/sh -c "echo '########################### - {}'" \; \
    -exec /bin/sh -c "{}" \;
fi

echo "###########################"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | finished entrypoint, starting vaulrdb-bin"

exec $@