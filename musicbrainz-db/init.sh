#!/bin/bash

set -o pipefail

SCRIPT_DIR="$(dirname $(realpath $BASH_SOURCE))"
PG_DSN=${MUSICBRAINZ_DSN:-'postgresql://musicbrainz:musicbrainz@localhost:15432/musicbrainz'}

echo "Creating schemata..."
psql $PG_DSN --csv -c 'CREATE SCHEMA IF NOT EXISTS musicbrainz'

echo "Installing extensions..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/Extensions.sql > /dev/null
echo "Creating collations..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateCollations.sql > /dev/null
echo "Creating tables..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateTables.sql > /dev/null

if [[ "$1" ]]
then
    echo "Loading data..."
    $SCRIPT_DIR/load.sh $1
else
    echo "No data path given. Skipping load."
fi

echo "Creating primary keys..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreatePrimaryKeys.sql > /dev/null
echo "Creating search configuration..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateSearchConfiguration.sql > /dev/null
echo "Creating functions..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateFunctions.sql > /dev/null
echo "Creating indexes..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateIndexes.sql > /dev/null
echo "Creating slave indexes..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateSlaveIndexes.sql > /dev/null
echo "Creating foreign key constraints..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateFKConstraints.sql > /dev/null
echo "Creating foreign table constraints..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateConstraints.sql > /dev/null
echo "Creating views..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateViews.sql > /dev/null
echo "Creating triggers..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateTriggers.sql > /dev/null
echo "Creating search indexes..."
psql $PG_DSN --csv -f $SCRIPT_DIR/mbdata/mbdata/sql/CreateSearchIndexes.sql > /dev/null
