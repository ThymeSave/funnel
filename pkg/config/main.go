package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// AppConfig represents the entire configuration for funnel
type AppConfig struct {
	Web *HTTP
}

// HTTP related configuration
type HTTP struct {
	Port        int      `env:"FUNNEL_PORT,default=3000"`
	CORSOrigins []string `env:"FUNNEL_CORS_ORIGINS,default=*"`
}

var config AppConfig

// ReadConfig from env variables
func ReadConfig(ctx context.Context) error {
	c := AppConfig{}
	err := envconfig.Process(ctx, &c)
	config = c
	return err
}

// Get current app configuration, make sure ReadConfig has been called before.
func Get() *AppConfig {
	return &config
}

// CreateDefault config and ignore errors, intended to be used by tests
func CreateDefault() {
	_ = ReadConfig(context.Background())
}
