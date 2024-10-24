package email

import (
	"context"

	"aming/go-nats/internal/models"

	uuid "github.com/satori/go.uuid"
)

// PGRepository Email postgresql repository interface
type PGRepository interface {
	Create(ctx context.Context, email *models.Email) (*models.Email, error)
	GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error)
	//Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.EmailsList, error)
}

// RedisRepository redis email repository interface
type RedisRepository interface {
	SetEmail(ctx context.Context, email *models.Email) error
	GetEmailByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error)
	DeleteEmail(ctx context.Context, emailID uuid.UUID) error
}
