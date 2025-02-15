package repository

import (
	"aming/go-nats/internal/models"
	"context"

	"aming/go-nats/pkg/utils"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type emailPGRepository struct {
	db *pgxpool.Pool
}

// NewEmailPGRepository Email postgresql constructor

func NewEmailPGRepository(db *pgxpool.Pool) *emailPGRepository {

	return &emailPGRepository{db: db}
}

// Create mail
func (e *emailPGRepository) Create(ctx context.Context, email *models.Email) (*models.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailPGRepository.Create")
	defer span.Finish()

	var mail models.Email
	if err := e.db.QueryRow(
		ctx,
		createEmailQuery,
		&email.From,
		&email.To,
		&email.Subject,
		&email.Message,
	).Scan(&mail.EmailID, &mail.From, &mail.To, &mail.Subject, &mail.Message, &mail.CreatedAt); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &mail, nil
}

// GetByID

// GetByID get single email by id
func (e *emailPGRepository) GetByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailPGRepository.GetByID")
	defer span.Finish()

	var mail models.Email
	if err := e.db.QueryRow(ctx, getByIDQuery, emailID).Scan(
		&mail.EmailID, &mail.From, &mail.To, &mail.Subject, &mail.Message, &mail.CreatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &mail, nil
}

// Search search email using postgresql full text search
func (e *emailPGRepository) Search(ctx context.Context, data string, pagination *utils.Pagination) (*models.EmailsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailPGRepository.Search")
	defer span.Finish()

	var count int
	if err := e.db.QueryRow(ctx, searchTotalCountQuery, data).Scan(&count); err != nil {
		return nil, errors.Wrap(err, "QueryRow")
	}
	if count == 0 {
		return &models.EmailsList{
			TotalCount: 0,
			TotalPages: 0,
			Page:       0,
			Size:       0,
			HasMore:    false,
			Emails:     make([]*models.Email, 0),
		}, nil
	}

	rows, err := e.db.Query(ctx, searchQuery, data, pagination.GetOffset(), pagination.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "db.Query")
	}
	defer rows.Close()

	emailList := make([]*models.Email, 0, count)
	for rows.Next() {
		var m models.Email
		if err := rows.Scan(&m.EmailID, &m.To, &m.From, &m.Subject, &m.Message, &m.CreatedAt); err != nil {
			return nil, errors.Wrap(err, " rows.Scan")
		}
		emailList = append(emailList, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows.Err")
	}

	return &models.EmailsList{
		TotalCount: int64(count),
		TotalPages: int64(pagination.GetTotalPages(count)),
		Page:       int64(pagination.GetPage()),
		Size:       int64(pagination.GetSize()),
		HasMore:    pagination.GetHasMore(count),
		Emails:     emailList,
	}, nil
}
