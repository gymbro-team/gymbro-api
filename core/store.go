package core

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataStore struct {
	db *sql.DB
}

var _store *DataStore

func InitializeStore(uri string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}

	_store = &DataStore{db}

	return nil
}

func DisconnectStore() error {
	return _store.db.Close()
}

func GetStore() *DataStore {
	return _store
}
