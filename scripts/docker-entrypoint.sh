#!/bin/sh
# docker-entrypoint.sh

# Immediately exit if any command has a non-zero exit status.
set -e

# Execute any pre-startup scripts or tasks.
/app/app db:migrate up

exec /app/app "$@"
