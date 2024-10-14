package middlewares

import (
	"aming/go-nats/config"
	"log"
)

type middlewareManager struct {
	cfg *config.Config
	log *log.Logger
}

func NewMiddlewareManager(log *log.Logger, cfg *config.Config) *middlewareManager {
	return &middlewareManager{log: log, cfg: cfg}
}
