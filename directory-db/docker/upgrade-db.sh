#!/bin/bash
set -e
set -o pipefail

sed -i "s/nlx-directory/${PGDATABASE}/g" /db-migrations/*.up.sql

/usr/local/bin/migrate \
    --database "postgres://${PGUSER}:${PGPASSWORD}@${PGHOST}:5432/${PGDATABASE}?sslmode=disable&connect_timeout=5" \
    --lock-timeout 600 \
    --path /db-migrations/ \
    up
