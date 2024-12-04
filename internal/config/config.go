package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DB Postgres

	Keys Keys

	Cache struct {
		Ttl int64 `mapstructure:"ttl"`
	} `mapstructure:"cache"`

	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Auth struct {
		TokenTTL        time.Duration `mapstructure:"token_ttl"`
		RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
	} `mapstructure:"auth"`
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Name     string
	SSLMode  string
	Password string
}

type Keys struct {
	Salt       string
	SigningKey string
}

func New(dirname, filename string) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(dirname)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("key", &cfg.Keys); err != nil {
		return nil, err
	}

	return cfg, nil
}
