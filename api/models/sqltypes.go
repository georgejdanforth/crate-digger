package models

import (
	"database/sql"
	"encoding/json"
)

type NullInt16 sql.NullInt16

func (ni *NullInt16) Scan(value interface{}) error {
	return (*sql.NullInt16)(ni).Scan(value)
}

func (ni *NullInt16) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int16)
}

type NullInt32 sql.NullInt32

func (ni *NullInt32) Scan(value interface{}) error {
	return (*sql.NullInt32)(ni).Scan(value)
}

func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}

type NullString sql.NullString

func (ns *NullString) Scan(value interface{}) error {
	return (*sql.NullString)(ns).Scan(value)
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}
