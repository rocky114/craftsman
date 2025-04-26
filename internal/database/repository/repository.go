package repository

import (
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sqlx.DB
}
