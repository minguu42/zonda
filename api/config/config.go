package config

import "time"

type Config struct {
	API API
}

type API struct {
	Host         string        `envconfig:"API_HOST" default:"0.0.0.0"`
	Port         int           `envconfig:"API_PORT" default:"8080"`
	ReadTimeout  time.Duration `envconfig:"API_READ_TIMEOUT" default:"2s"`
	WriteTimeout time.Duration `envconfig:"API_WRITE_TIMEOUT" default:"2s"`
	StopTimeout  time.Duration `envconfig:"API_STOP_TIMEOUT" default:"25s"`
}
