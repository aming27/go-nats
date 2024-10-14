package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

//Configuration microservice

type Config struct {
	AppVersion string
	HTTP       HTTP
	GRPC       GRPC
	Logger     Logger
	//Metrics    Metrics
	//Jaeger      Jaeger
	Nats  Nats
	Redis Redis
	//MailService MailService
	PostgreSQL PostgreSQL
}

// HTTP server configuration
type HTTP struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

// Logger configuration
type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Redis Configuration
type Redis struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultDB string
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    int
	DB             int
}

// Nats config
type Nats struct {
	URL       string
	ClusterID string
	ClientID  string
}

// PostgreSQL config
type PostgreSQL struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDBName   string
	PostgresqlSSLMode  string
	PgDriver           string
}

// GRPC gRPC service config
type GRPC struct {
	Port              string
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
}

func exportConfig() error {

	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// ParseConfig Parse config file
func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil

}
