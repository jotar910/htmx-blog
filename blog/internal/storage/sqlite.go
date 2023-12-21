package storage

import "github.com/jmoiron/sqlx"

type SQLiteDatabase struct {
	db *sqlx.DB
}

func NewSQLiteDatabase(db *sqlx.DB) *SQLiteDatabase {
	return &SQLiteDatabase{db: db}
}
