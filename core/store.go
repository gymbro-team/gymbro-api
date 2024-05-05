package core

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

var db *Database

func InitializeDatabase(dataSourceName string) {
	database, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		panic(err)
	}

	db = &Database{database}
}

func GetDB() *sql.DB {
	return db.db
}

func CloseDB() {
	db.db.Close()
}
