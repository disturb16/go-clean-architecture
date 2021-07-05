package sqlite

import (
	"context"

	"github.com/disturb16/go-sqlite-service/internal/persons/repository/rediscache"
	"github.com/jmoiron/sqlx"
)

type Sqlite struct {
	db    *sqlx.DB
	cache *rediscache.Cache
}

func New(db *sqlx.DB, redisCache *rediscache.Cache) (sl *Sqlite) {
	return &Sqlite{
		db:    db,
		cache: redisCache,
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
