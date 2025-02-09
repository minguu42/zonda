package clocktest

import (
	"context"
	"testing"
	"time"

	"github.com/minguu42/zonda/lib/go/clock/internal"
)

type nowKey struct{}

func init() {
	if testing.Testing() {
		internal.Now = nowForTest
	}
}

func nowForTest(ctx context.Context) time.Time {
	if now, ok := ctx.Value(nowKey{}).(time.Time); ok {
		return now
	}
	return internal.DefaultNow(ctx)
}

func WithFixedNow(t *testing.T, ctx context.Context, tm time.Time) context.Context {
	t.Helper()
	return context.WithValue(ctx, nowKey{}, tm)
}
