package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func Load() (*Config, error) {
	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		return nil, fmt.Errorf("failed to populate the specified struct based on environment variables: %w", err)
	}
	return &conf, nil
}
