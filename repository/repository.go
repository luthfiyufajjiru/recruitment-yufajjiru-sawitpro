// This file contains the repository implementation layer.
package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	IntegrityViolationClassCode = "23"
)

type Repository struct {
	Db        *sqlx.DB
	SaltSize  int
	SecretKey []byte
}

type NewRepositoryOptions struct {
	Dsn       string
	SaltSize  int
	SecretKey string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sqlx.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db:        db,
		SaltSize:  opts.SaltSize,
		SecretKey: []byte(opts.SecretKey),
	}
}
