package server

import (
	"aming/go-nats/config"
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type server struct {
	cfg   *config.Config
	redis *redis.Client
	echo  *echo.Echo
	nats  nats.JetStreamContext
}

func NewServer(
	cfg *config.Config,
	redis *redis.Client,
	nats nats.JetStreamContext,

) *server {

	return &server{cfg: cfg, redis: redis, nats: nats, echo: echo.New()}
}

func (s *server) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Printf("Server is listening on PORT: %s", s.cfg.HTTP.Port)
		s.runHttpServer()
	}()

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "echo.Server.Shutdown")

	}

	return nil

}
