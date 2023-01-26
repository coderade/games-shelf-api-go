package db

import "database/sql"

type DBHelper struct {
	DB *sql.DB
}

// NewDBHelper creates a new instance of DBHelper
func NewDBHelper(db *sql.DB) *DBHelper {
	return &DBHelper{DB: db}
}

// Close closes the database connection
func (helper *DBHelper) Close() error {
	return helper.DB.Close()
}
