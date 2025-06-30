package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Listener string `envconfig:"LISTENER" default:":8080"`
	Host     string `envconfig:"PUBLIC_HOSTNAME" default:"http://localhost:8080"`

	RedisUrl  string        `envconfig:"REDIS_URL" required:"true"`
	CacheTTL  time.Duration `envconfig:"CACHE_TTL" required:"true"`
	JWTSecret string        `envconfig:"JWT_SECRET" required:"true"`
}

var Data Config

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Debug("cannot load .env file", slog.Any("error", err))
	}
	if err := envconfig.Process("FILE_CACHER", &Data); err != nil {
		slog.Error("cannot read configuration", slog.Any("error", err))
		os.Exit(1)
	}
}

func configError(err error) {
	slog.Error("Failed to load config", "error", err)
	os.Exit(1)
}
