package config

import (
	"time"

	"github.com/leoomi/remnants/pkg/cfgloader"
)

type Config struct {
	AuthPort            string        `mapstructure:"AUTH_PORT"`
	AuthServerAddress   string        `mapstructure:"AUTH_SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (*Config, error) {
	return cfgloader.LoadConfig[Config](path)
}
