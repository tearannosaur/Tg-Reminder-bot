package handlers

import (
	"reminder/repository"
)

type HandlerModule struct {
	repo *repository.ReposytoryModule
}

func HandlerModuleInit(r *repository.ReposytoryModule) *HandlerModule {
	return &HandlerModule{
		repo: r,
	}
}
