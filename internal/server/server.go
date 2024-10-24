package server

import (
	"aming/go-nats/config"
	natsPublisher "aming/go-nats/internal/email/delivery/nats"
	"aming/go-nats/internal/email/repository"
	"aming/go-nats/internal/email/usecase"
	"aming/go-nats/smtp"
	"context"
	"log"
	"net/smtp"
	"syscall"

	"os"
	"os/signal"

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
	nats    nats.Conn
	pgxPool *pgxpool.Pool
	log     *log.Logger
	smtp    *smtp.Client
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
	publisher, _ := natsPublisher.NewPublisher(s.cfg.Nats.URL)
	emailPgRepo := repository.NewEmailPGRepository(s.pgxPool)
	emailRedisRepo := repository.NewEmailRedisRepository(s.redis)
	emailUC := usecase.NewEmailUseCase(s.log, emailPgRepo, publisher, smtpClient, emailRedisRepo)

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
