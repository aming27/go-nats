package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type emailPGRepository struct {
	db *pgxpool.Pool
}

// NewEmailPGRepository Email postgresql constructor

func NewEmailPGRepository(db *pgxpool.Pool) *emailPGRepository {

	return &emailPGRepository{db: db}
}
