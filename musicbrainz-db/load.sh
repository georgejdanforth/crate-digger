#!/bin/bash

set -o pipefail

PG_DSN=${MUSICBRAINZ_DSN:-'postgresql://musicbrainz:musicbrainz@localhost:15432/musicbrainz'}
DATA_DIR=""


if [[ "$1" ]]
then
    DATA_DIR=$1
else
    echo "Need directory of mbdump"
    exit 1
fi

if [ ! -d $DATA_DIR ]
then
    echo "$DATA_DIR is not a directory"
    exit 1
fi

function import_file {
    echo "$1"
    echo "$(basename $1)"
}
export -f import_file

for path in `find $DATA_DIR | tail -n +2`
do
    table_name=$(basename $path)
    # cmd="COPY $table_name FROM S"
    # echo "psql $PG_DSN -c '$cmd'"
    cat $path | psql $PG_DSN -c "COPY $table_name FROM STDIN"
done
