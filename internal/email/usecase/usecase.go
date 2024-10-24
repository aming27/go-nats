package usecase

import (
	"aming/go-nats/internal/email"
	"aming/go-nats/internal/email/delivery/nats"
	smtpClient "aming/go-nats/pkg/smtp"

	"github.com/nats-io/gnatsd/logger"
)

const (
	createEmailSubject = "mail:create"
	sendEmailSubject   = "mail:send"
)

type emailUseCase struct {
	log         logger.Logger
	emailPGRepo email.PGRepository
	publisher   nats.Publisher
	smtpClient  smtpClient.SMTPClient
	redisRepo   email.RedisRepository
}

// NewEmailUseCase email usecase constructor
func NewEmailUseCase(log logger.Logger, emailPGRepo email.PGRepository, publisher nats.Publisher, smtpClient smtpClient.SMTPClient, redisRepo email.RedisRepository) *emailUseCase {
	return &emailUseCase{log: log, emailPGRepo: emailPGRepo, publisher: publisher, smtpClient: smtpClient, redisRepo: redisRepo}
}
