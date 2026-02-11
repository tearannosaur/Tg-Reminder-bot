package repository

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryModule struct {
	db *sqlx.DB
}

func ModuleInit(d *sqlx.DB) *RepositoryModule {
	return &RepositoryModule{
		db: d,
	}
}
