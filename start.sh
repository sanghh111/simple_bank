#!/bin/sh

set -e
echo "run db migration $URI_DB"
/app/migrate --path /app/migration --database $URI_DB --verbose up

echo "run server"
exec $@