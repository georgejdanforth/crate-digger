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
