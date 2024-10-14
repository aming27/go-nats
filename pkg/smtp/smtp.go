package smtp

import (
	"aming/go-nats/config"
	"aming/go-nats/internal/models"
)

// SMTPClient interface
type SMTPClient interface {
	SendMail(mail *models.MailData) error
}

type smtpClient struct {
	cfg *config.Config
}

// NewSmtpClient constructor
func NewSmtpClient(cfg *config.Config) *smtpClient {
	return &smtpClient{cfg: cfg}
}
