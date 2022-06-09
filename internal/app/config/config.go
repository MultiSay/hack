package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host            string `envconfig:"APP_HOST" default:"localhost"`
	Port            int    `envconfig:"APP_PORT" default:"8080"`
	URL             string `envconfig:"DATABASE_URL" default:"postgresql://user:password@database:5432/hack?sslmode=disable"`
	MaxOpenConns    int    `envconfig:"MAX_OPEN_CONNS" default:"25"`
	MaxIdleConns    int    `envconfig:"MAX_IDLE_CONNS" default:"2"`
	ConnMaxLifetime int    `envconfig:"CONN_MAX_LIFETIME" default:"1"`
	DeliveryMaxTime int    `envconfig:"DELIVERY_MAX_TIME" default:"30"`
	SigningKey      string `envconfig:"SIGNING_KEY" default:"some_secret_key"`
}

var (
	cfg *Config
	mx  sync.RWMutex
)

func initConfig() (*Config, error) {
	godotenv.Load()
	cfg = &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Set main config
func Set(c Config) {
	mx.Lock()
	defer mx.Unlock()
	cfg = &c
}

// Get main config
func Get() Config {
	var err error
	mx.RLock()
	defer mx.RUnlock()
	if cfg == nil {
		cfg, err = initConfig()
		if err != nil {
			panic(err)
		}
	}
	return *cfg
}

// Reload main config
func Reload() Config {
	cfg, err := initConfig()
	if err != nil {
		return Get()
	}
	Set(*cfg)
	return Get()
}
