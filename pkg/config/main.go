package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// AppConfig represents the entire configuration for funnel
type AppConfig struct {
	// Web related configuration
	Web *HTTP
	// CouchDB related configuration
	CouchDB *CouchDB
}

// HTTP related configuration
type HTTP struct {
	// Port to use
	Port int `env:"FUNNEL_PORT,default=3000"`
	// CORSOrigins to allow
	CORSOrigins []string `env:"FUNNEL_CORS_ORIGINS,default=*"`
}

// CouchDB related configuration
type CouchDB struct {
	// Scheme for the http calls
	Scheme string `env:"FUNNEL_COUCHDB_SCHEME,default=http"`
	// Host for the http calls
	Host string `env:"FUNNEL_COUCHDB_HOST,default=127.0.0.1"`
	// Port for the http calls
	Port int `env:"FUNNEL_COUCHDB_PORT,default=5984"`
	// AdminUser is the username of the administrator to use for user creation and operative tasks, it must already exist
	AdminUser string `evn:"FUNNEL_COUCHDB_ADMIN_USER,default=admin"`
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
