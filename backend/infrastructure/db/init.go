package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct{}
func (db DB)DbConn() *bun.DB {
	dsn := "postgres://postgres:@db:5432/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn), pgdriver.WithPassword("postgres")))
	return bun.NewDB(sqldb, pgdialect.New())
}