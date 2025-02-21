package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config
	Config struct {
		HTTP
		PG
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" env:"WORKING_PORT" env-upd:"true"`
	}
	// PG -.
	PG struct {
		Host     string `env-required:"true" env:"DB_HOST"`
		Port     string `env-required:"true" env:"DB_PORT"`
		Database string `env-required:"true" env:"DB_DATABASE"`
		Username string `env-required:"true" env:"DB_USERNAME"`
		Password string `env-required:"true" env:"DB_PASSWORD"`
		SslMode  string `env-required:"true" env:"DB_SSL_MODE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		errMessage := err.Error()
		fmt.Printf("Environment variable error: %s, trying to read from .env\n", errMessage)
		err = cleanenv.ReadConfig(".env", cfg)
		if err != nil {
			if os.IsNotExist(err) {
				// when working directory is src/internal/tests
				err = cleanenv.ReadConfig(filepath.Join("..", ".env"), cfg)
			}
			if err != nil {
				return nil, err
			}
		}
		fmt.Print("Environment variable success: variables successfully read from .env\n")
	}

	return cfg, nil
}

func (cfg *Config) PGConnectionURLString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?timezone=%s&sslmode=%s",
		cfg.PG.Username,
		cfg.PG.Password,
		cfg.PG.Host,
		cfg.PG.Port,
		cfg.PG.Database,
		"UTC",
		cfg.PG.SslMode,
	)
}
