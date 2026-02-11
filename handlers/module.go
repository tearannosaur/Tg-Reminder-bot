package handlers

import (
	"reminder/repository"
)

type HandlerModule struct {
	repo *repository.RepositoryModule
}

func HandlerModuleInit(r *repository.RepositoryModule) *HandlerModule {
	return &HandlerModule{
		repo: r,
	}
}
