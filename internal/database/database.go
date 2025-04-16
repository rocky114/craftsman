package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database/sqlc"
	"time"
)

type Store struct {
	db *sql.DB
	*sqlc.Queries
}

func NewStore(cfg config.DatabaseConfig) (*Store, error) {
	db, err := sql.Open("mysql", cfg.URL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Hour)

	return &Store{
		db:      db,
		Queries: sqlc.New(db),
	}, nil
}
