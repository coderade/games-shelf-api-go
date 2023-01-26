package db

import (
	"context"
	"database/sql"
)

type Database interface {
	PingContext(ctx context.Context) error
	Close() error
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type SQLDatabase struct {
	DB *sql.DB
}

func (d *SQLDatabase) PingContext(ctx context.Context) error {
	return d.DB.PingContext(ctx)
}

func (d *SQLDatabase) Close() error {
	return d.DB.Close()
}

func (d *SQLDatabase) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.DB.QueryContext(ctx, query, args...)
}

func (d *SQLDatabase) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRowContext(ctx, query, args...)
}

func (d *SQLDatabase) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.DB.ExecContext(ctx, query, args...)
}
