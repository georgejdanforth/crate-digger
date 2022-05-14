#!/bin/bash

set -o pipefail

SCRIPT_DIR="$(dirname $(realpath $BASH_SOURCE))"
PG_DSN=${MUSICBRAINZ_DSN:-'postgresql://musicbrainz:musicbrainz@localhost:15432/musicbrainz'}

echo "Creating schemata..."
psql $PG_DSN --csv -c 'CREATE SCHEMA IF NOT EXISTS musicbrainz'

echo "Instaling Extensions..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/Extensions.sql
echo "Creating collations..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateCollations.sql
echo "Creating types..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateTypes.sql
echo "Creating tables..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateTables.sql
echo "Creating primary keys..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreatePrimaryKeys.sql
echo "Creating search configuration..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateSearchConfiguration.sql
echo "Creating functions..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateFunctions.sql
echo "Creating indexes..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateIndexes.sql
echo "Creating foreign key constraints..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateFKConstraints.sql
echo "Creating foreign table constraints..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateConstraints.sql
echo "Setting raw initial sequence values..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/SetSequences.sql
echo "Creating views..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateViews.sql
echo "Creating triggers..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateTriggers.sql
echo "Creating search indexes..."
psql $PG_DSN --csv -f $SCRIPT_DIR/musicbrainz-server/admin/sql/CreateSearchIndexes.sql
