package database

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database/repository"
	"github.com/rocky114/craftman/internal/database/sqlc"
	"time"
)

type Database struct {
	*sqlx.DB
	*sqlc.Queries
	*repository.Repository
}

func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
	db, err := sqlx.Open("mysql", cfg.URL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Hour)

	return &Database{
		DB:         db,
		Queries:    sqlc.New(db),
		Repository: repository.NewRepository(db),
	}, nil
}

type TxFunc func(queries *sqlc.Queries) error

func (s *Database) WithTransaction(ctx context.Context, fn TxFunc) error {
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
