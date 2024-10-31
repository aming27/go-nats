package email

import (
	"aming/go-nats/internal/models"
	"aming/go-nats/pkg/utils"
	"context"

	uuid "github.com/satori/go.uuid"
)

// UseCase Email usecase interface
type UseCase interface {
	Create(ctx context.Context, email *models.Email) error
	PublishCreate(ctx context.Context, email *models.Email) error
	GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error)
	Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.EmailsList, error)
	SendEmail(ctx context.Context, email *models.Email) error
}
