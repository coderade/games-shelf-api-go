package models

import "database/sql"

type DBHelper struct {
	DB *sql.DB
}
