package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() (*Database, error) {
	db, err := sqlx.Connect("mysql", "root:@/danil")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}
