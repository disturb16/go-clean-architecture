package repository

import (
	"context"
	"errors"

	"github.com/disturb16/go-clean-architecture/internal/persons"
	"github.com/disturb16/go-clean-architecture/internal/persons/repository/mysql"
	"github.com/disturb16/go-clean-architecture/internal/persons/repository/sqlite"
	"github.com/disturb16/go-clean-architecture/settings"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

// New constructs the repository
func New(ctx context.Context, cfg *settings.Settings, db *sqlx.DB) (persons.Repository, error) {

	switch cfg.DB.Engine {
	case "sqlite":
		err := sqlite.CreateSchema(ctx, db)
		if err != nil {
			return nil, err
		}
		return sqlite.New(db), nil

	case "mysql":
		return mysql.New(db), nil

	default:
		err := errors.New("unsupported or missing database engine")
		return nil, err
	}
}
