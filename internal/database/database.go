package database

import (
	"database/sql"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database/sqlc"
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

	return &Store{
		db:      db,
		Queries: sqlc.New(db),
	}, nil
}
