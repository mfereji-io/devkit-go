package contexthelper

import (
	"context"

	"github.com/google/uuid"
)

func WithAppId(ctx context.Context, appId string) context.Context {
	return context.WithValue(ctx, ContextKeyAppId, appId)
}

func GetRequestAppId(ctx context.Context) string {

	existing := ctx.Value(ContextKeyAppId)

	if existing != nil {

		if val, ok := existing.(string); ok {

			if u, err := uuid.Parse(val); err == nil {
				return u.String()
			}
			return ""
		}
		return ""
	}
	return ""
}
