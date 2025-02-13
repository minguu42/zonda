package config_test

import (
	"testing"
	"time"

	"github.com/minguu42/zonda/api/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	want := config.Config{
		API: config.API{
			Host:         "0.0.0.0",
			Port:         8080,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
			StopTimeout:  25 * time.Second,
		},
	}
	got, err := config.Load()
	require.NoError(t, err)
	assert.Equal(t, want, got)
}
