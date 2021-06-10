package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Sqlite struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) (sl *Sqlite) {
	return &Sqlite{
		db: db,
	}
}

func CreateSchema(ctx context.Context, db *sqlx.DB) error {
	createPersonsTable := `
		CREATE TABLE IF NOT EXISTS persons(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INTEGER
		)`

	_, err := db.ExecContext(ctx, createPersonsTable)
	if err != nil {
		return err
	}

	return nil
}

func (sl *Sqlite) Close() error {
	return sl.db.Close()
}
