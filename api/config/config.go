package config

import "time"

type Config struct {
	API  API
	Auth Auth
	DB   DB
}

type API struct {
	Host         string        `env:"API_HOST" default:"0.0.0.0"`
	Port         int           `env:"API_PORT" default:"8080"`
	ReadTimeout  time.Duration `env:"API_READ_TIMEOUT" default:"2s"`
	WriteTimeout time.Duration `env:"API_WRITE_TIMEOUT" default:"2s"`
	StopTimeout  time.Duration `env:"API_STOP_TIMEOUT" default:"25s"`
}

type Auth struct {
	AccessTokenExpiry  time.Duration `env:"ACCESS_TOKEN_EXPIRY" default:"2h"`
	RefreshTokenExpiry time.Duration `env:"REFRESH_TOKEN_EXPIRY" default:"168h"`
	AccessTokenSecret  string        `env:"ACCESS_TOKEN_SECRET,required"`
	RefreshTokenSecret string        `env:"REFRESH_TOKEN_SECRET,required"`
}

type DB struct {
	Host            string        `env:"DB_HOST,required"`
	Port            int           `env:"DB_PORT,required"`
	Database        string        `env:"DB_DATABASE,required"`
	User            string        `env:"DB_USER,required"`
	Password        string        `env:"DB_PASSWORD,required"`
	MaxOpenConns    int           `env:"DB_MAX_OPEN_CONNS" default:"25"`
	MaxIdleConns    int           `env:"DB_MAX_IDLE_CONNS" default:"25"`
	ConnMaxLifetime time.Duration `env:"DB_CONN_MAX_LIFETIME" default:"5m"`
}
