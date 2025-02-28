package config

import "time"

type Config struct {
	API API
}

type API struct {
	Host         string        `env:"API_HOST" default:"0.0.0.0"`
	Port         int           `env:"API_PORT" default:"8080"`
	ReadTimeout  time.Duration `env:"API_READ_TIMEOUT" default:"2s"`
	WriteTimeout time.Duration `env:"API_WRITE_TIMEOUT" default:"2s"`
	StopTimeout  time.Duration `env:"API_STOP_TIMEOUT" default:"25s"`
}
