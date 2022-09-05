package config

import "time"

type App struct {
	LogLevel string `default:"debug"`
	JWT      JWT
}

type JWT struct {
	TokenTTLHours time.Duration `default:"1h"`
	SecretKey     string        `default:"secret"` // *** IMPORTANT In production must be configured as a secret, not as env variable
}
