package dbutils

import (
	"errors"
	"fmt"
	"log"

	"github.com/disturb16/go-sqlite-service/settings"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// ErrInvalidDBEngine database engine is not supported
	ErrInvalidDBEngine = errors.New("uUnsoported or missing database engine")
)

func New(config *settings.Settings) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	dbConfig := config.DB

	switch dbConfig.Engine {
	case "mysql":
		db, err = CreateMySqlConnection(config)

	case "sqlite":
		db, err = CreateSqliteConnection(config)

	default:
		err = ErrInvalidDBEngine
	}

	return db, err
}

func CreateMySqlConnection(config *settings.Settings) (*sqlx.DB, error) {

	var connectionString string
	var db *sqlx.DB
	var err error
	dbConfig := config.DB

	connectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	log.Println("Trying to connect to database...")
	db, err = sqlx.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}

func CreateSqliteConnection(config *settings.Settings) (*sqlx.DB, error) {
	source := fmt.Sprintf("./%s.db", config.DB.Name)
	db, err := sqlx.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
