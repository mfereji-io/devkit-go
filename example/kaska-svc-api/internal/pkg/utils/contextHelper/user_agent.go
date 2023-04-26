package contexthelper

import (
	"context"
)

func UserAgent(ctx context.Context) string {
	existing := ctx.Value(ContextKeyUserAgent)
	if existing == nil {
		return ""
	}

	return existing.(string)
}

func WithUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, ContextKeyUserAgent, userAgent)
}
