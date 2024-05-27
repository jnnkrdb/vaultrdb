#!/bin/sh

set -e

# prepare the index.html for the frontend

echo "$(date +"%Y-%m-%d - %H:%M:%S") | prepare index.html"

web_dir="$VRDB_DIRECTORY_ROOT/web"
_host="$(hostname)"

echo "$(date +"%Y-%m-%d - %H:%M:%S") | configs:"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | -- web directory: $web_dir"
echo "$(date +"%Y-%m-%d - %H:%M:%S") | ------- hostname: $_host"

# replace the hostname in the index.html
sed -i "s|{{HOSTNAME}}|$_host|g" $web_dir/index.html 
 
echo "$(date +"%Y-%m-%d - %H:%M:%S") | index.html prepared"