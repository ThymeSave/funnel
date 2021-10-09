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
	// Oauth2 related configuration
	Oauth2 *OAuth2
}

// OAuth2 represents the configuration related to oauth authentication
type OAuth2 struct {
	// IssuerURL is the base url where to find the endpoint `/.well-known/openid-configuration` with
	// configuration for oidc
	IssuerURL string `env:"FUNNEL_OAUTH2_ISSUER_URL,required"`
	// ClientID for jwt tokens to validate
	ClientID string `env:"FUNNEL_OAUTH2_CLIENT_ID,required"`
	// VerifyIssuer in jwt claims
	VerifyIssuer bool `env:"FUNNEL_OAUTH2_VERIFY_ISSUER,default=true"`
	// UsernameClaim is the name of the claim that will be used to uniquely identify the user
	UsernameClaim string `env:"FUNNEL_OAUTH2_USERNAME_CLAIM,default=email"`
	// Scopes to include into new JWTs
	Scopes []string `env:"FUNNEL_OAUTH2_SCOPES,default=openid,profile,email"`
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
	Get().Oauth2.VerifyIssuer = false
}
