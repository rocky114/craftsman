package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database/sqlc"
	"time"
)

type Repository struct {
	*sql.DB
	*sqlc.Queries
}

func NewRepository(cfg config.DatabaseConfig) (*Repository, error) {
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

	return &Repository{
		DB:      db,
		Queries: sqlc.New(db),
	}, nil
}

type TxFunc func(queries *sqlc.Queries) error

func (s *Repository) WithTransaction(ctx context.Context, fn TxFunc) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		}
	}()

	txQueries := s.Queries.WithTx(tx)

	if err = fn(txQueries); err != nil {
		return fmt.Errorf("execute transaction: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	return nil
}
