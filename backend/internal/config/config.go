package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
	
	"github.com/joho/godotenv"
)

type Config struct {
	App				AppConfig
	DB				DBConfig
	Google  		GoogleConfig
	Session 		SessionConfig
}

type AppConfig struct {
	Env				string
	Port    		string
	Scheme  		string
}

type DBConfig struct {
	URL				string
}

type GoogleConfig struct {
	ClientID		string
	ClientSecret	string
	RedirectURL 	string
}

type SessionConfig struct {
	Duration		time.Duration
	CookieDomain	string
	CookieSecure	bool
}

// Load loads the configuration from the environment var
func Load() (*Config, error) {
	_ = godotenv.Load()

	sessionHours, err := strconv.Atoi(getEnv("SESSION_DURATION_HOURS", "720"))
	if err != nil {
		return nil, fmt.Errorf("Invalid SESSION_DURATION_HOURS: %w", err)
	}

	cookieSecure, _ := strconv.ParseBool(getEnv("SESSION_COOKIE_SECURE", "false"))

	cfg := &Config {
		App: AppConfig {
			Env: getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
			Scheme: getEnv("APP_SCHEME", "popout://"),
		},
		DB: DBConfig {
			URL: requireEnv("DATABASE_URL"),
		},
		Google: GoogleConfig {
			ClientID: requireEnv("GOOGLE_CLIENT_ID"),
			ClientSecret: requireEnv("GOOGLE_CLIENT_SECRET"),
			RedirectURL: requireEnv("GOOGLE_REDIRECT_URL"),
		},
		Session: SessionConfig {
			Duration: time.Duration(sessionHours) * time.Hour,
			CookieDomain: getEnv("SESSION_COOKIE_DOMAIN", "localhost"),
			CookieSecure: cookieSecure,
		},
	}

	return cfg, nil

}

// IsProd check if is on prod env
func (c *Config) IsProd() bool {
	return c.App.Env == "production"
}

// getEnv gets the environment var with a default val
func getEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return defaultValue
}

// requireEnv gets the environment var and panics if it is not set
func requireEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("[CONFIG] Required environment variable %q is not set", key)) 
	}

	return value
}