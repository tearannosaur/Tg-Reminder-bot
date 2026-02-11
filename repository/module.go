package repository

import (
	"github.com/jmoiron/sqlx"
)

type ReposytoryModule struct {
	db *sqlx.DB
}

func ModuleInit(d *sqlx.DB) *ReposytoryModule {
	return &ReposytoryModule{
		db: d,
	}
}
