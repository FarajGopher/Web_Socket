package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func InitDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:Far@1234@localhost:5432/chat_group?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
