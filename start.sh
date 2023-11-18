#!/bin/sh

set -e
echo "run db migration"
/app/migrate --path /app/migration --database $URI_DB --verbose up

echo "run server $@"
exec "$@"