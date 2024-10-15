package postgresql

import (
	"aming/go-nats/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifeTime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

//NewpgConn Pool

func NewPgxConn(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dataSourceConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.PostgreSQL.PostgresqlHost,
		cfg.PostgreSQL.PostgresqlPort,
		cfg.PostgreSQL.PostgresqlUser,
		cfg.PostgreSQL.PostgresqlDBName,
		cfg.PostgreSQL.PostgresqlSSLMode,
		cfg.PostgreSQL.PostgresqlPassword,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceConfig)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifeTime
	poolCfg.MinConns = minConns
	poolCfg.LazyConnect = lazyConnect

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)

	if err != nil {
		return nil, err
	}
	return connPool, err

}
