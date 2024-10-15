package main

import (
	"aming/go-nats/config"
	"aming/go-nats/internal/server"
	"aming/go-nats/pkg/nats"
	"aming/go-nats/pkg/postgresql"
	"aming/go-nats/pkg/redis"

	"log"
)

func main() {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		log.Printf("NewRedisClient: %+v\n", err)
	}
	log.Printf("Redis connected: %+v\n", redisClient.PoolStats())

	nats.JetStreamInit()
	if err != nil {
		log.Printf("JetStreamInit: %+v\n", err)
	}

	pgxPool, err := postgresql.NewPgxConn(cfg)
	if err != nil {
		log.Printf("NewPgxConn: %+v\n", err)
	}
	log.Printf("PostgreSQL connected: %+v", pgxPool.Stat().TotalConns())

	s := server.NewServer(cfg, redisClient, pgxPool)
	s.Run()
}

// Connect to Redis
// redisClient, err := redis.NewRedisClient(cfg)
// if err != nil {
// 	log.Fatal(err)
// }

// 	print(cfg)
// 	log.Printf("HTTP Port: %s\n", cfg.AppVersion)

// }
// func print(l *config.Config) {
// 	log.Printf("HTTP Port: %s\n", l.AppVersion)
// }
