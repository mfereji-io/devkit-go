package contexthelper

import (
	"context"
)

func WithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestId, requestId)
}

/*
func getRequestId(ctx context.Context) string {

	existing := ctx.Value(ContextKeyRequestId)
	if existing == nil {
		return ""
	}

	if val, ok := existing.(string); ok {
		u, err := uuid.Parse(val)
		if err != nil {
			return val
		}
		return u.String()
	}

	return ""
}
*/
