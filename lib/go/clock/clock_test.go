package clock_test

import (
	"context"
	"testing"
	"time"

	"github.com/minguu42/zonda/lib/go/clock"
	"github.com/minguu42/zonda/lib/go/clock/clocktest"
	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	want := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	ctx := clocktest.WithFixedNow(t, context.Background(), want)

	got := clock.Now(ctx)
	assert.Equal(t, want, got)
}
