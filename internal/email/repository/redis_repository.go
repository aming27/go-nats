package repository

import (
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	prefix     = "emails"
	expiration = time.Second * 3600
)

type emailRedisRepository struct {
	redis *redis.Client
}

// NewEmailRedisRepository emails redis repository constructor
func NewEmailRedisRepository(redis *redis.Client) *emailRedisRepository {
	return &emailRedisRepository{redis: redis}
}
