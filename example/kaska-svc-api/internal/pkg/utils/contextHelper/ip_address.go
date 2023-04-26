package contexthelper

import (
	"context"
)

func IPAddress(ctx context.Context) string {
	existing := ctx.Value(ContextKeyIPAddress)
	if existing == nil {
		return ""
	}

	return existing.(string)
}

func WithIpAddress(ctx context.Context, ipAddress string) context.Context {
	return context.WithValue(ctx, ContextKeyIPAddress, ipAddress)
}
