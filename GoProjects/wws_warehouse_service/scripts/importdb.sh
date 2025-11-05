#!/bin/bash

# Usage: Ensure environment variables DB_USER and DB_PASS are set.
# Optional: set DB_NAME (defaults to "database_name").
# This script imports all .sql files from ./.database/sample into the target MySQL database.

set -euo pipefail

SQL_DIR="./database/data_samples"
DB_USER="${DB_USER:?Environment variable DB_USER is required}"
DB_PASS="${DB_PASS:?Environment variable DB_PASS is required}"
DB_NAME="${DB_NAME:-database_name}"

if [ ! -d "$SQL_DIR" ]; then
  echo "Error: SQL directory '$SQL_DIR' does not exist." >&2
  exit 1
fi

shopt -s nullglob
SQL_FILES=("$SQL_DIR"/*.sql)
shopt -u nullglob

if [ ${#SQL_FILES[@]} -eq 0 ]; then
  echo "No .sql files found in $SQL_DIR"
  exit 0
fi

for f in "${SQL_FILES[@]}"; do
  echo "Importing $f into database '$DB_NAME'..."
  mysql -h 127.0.0.1 -u"$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$f"
  echo "Imported $f"
done
