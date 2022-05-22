package database

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() {
	// TODO: make connection params configurable
	db = sqlx.MustConnect(
		"postgres", "host=localhost port=15432 user=musicbrainz password=musicbrainz dbname=musicbrainz sslmode=disable")
}

func GetDb() *sqlx.DB {
	return db
}
