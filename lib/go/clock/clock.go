package clock

import (
	"context"
	"time"

	"github.com/minguu42/zonda/lib/go/clock/internal"
)

func Now(ctx context.Context) time.Time {
	return internal.Now(ctx)
}
