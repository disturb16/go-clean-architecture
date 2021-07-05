package mysql

import (
	"github.com/disturb16/go-sqlite-service/internal/persons/repository/rediscache"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
)

// Mysql connection
type Mysql struct {
	db    *sqlx.DB
	cache *rediscache.Cache
}

// New returns an instance of Mysql connection
func New(db *sqlx.DB, redisCache *rediscache.Cache) (m *Mysql) {
	return &Mysql{
		db:    db,
		cache: redisCache,
	}
}

func (m *Mysql) Close() error {
	return m.db.Close()
}
