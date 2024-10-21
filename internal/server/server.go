package server

import (
	"aming/go-nats/config"
	"aming/go-nats/pkg/smtp"
	"context"
	"log"
	"syscall"

	"os"
	"os/signal"

	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type server struct {
	cfg     *config.Config
	redis   *redis.Client
	echo    *echo.Echo
	nats    nats.JetStreamContext
	pgxPool *pgxpool.Pool
	log     *log.Logger
}

func NewServer(
	cfg *config.Config,
	redis *redis.Client,
	pgxPool *pgxpool.Pool,
	//nats nats.JetStreamContext,

) *server {

	return &server{cfg: cfg, redis: redis, pgxPool: pgxPool, echo: echo.New()}
}

func (s *server) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	smtpClient := smtp.NewSmtpClient(s.cfg)

	//mv := middlewares.NewMiddlewareManager(s.log, s.cfg)

	//validate := validator.New()

	go func() {
		log.Printf("Server is listening on PORT: %s", s.cfg.HTTP.Port)
		s.runHttpServer()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatal("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatal("ctx.Done: %v", done)
	}

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "echo.Server.Shutdown")

	}

	return nil

}
