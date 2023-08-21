package handler

import (
	"pathfinder-family/config"
	"pathfinder-family/data/cache/cacheInterface"
	"pathfinder-family/data/db/dbInterface"
)

type Handler struct {
	cfg      *config.Config
	postgres dbInterface.IPostgres
	inmemory cacheInterface.IInMemory
}

func NewHandler(cfg *config.Config, postgres dbInterface.IPostgres, inmemory cacheInterface.IInMemory) *Handler {
	return &Handler{
		cfg:      cfg,
		postgres: postgres,
		inmemory: inmemory,
	}
}
