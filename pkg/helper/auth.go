package helper

import (
	"context"
	"todo-api-gin-gorm/pkg/config"
)

// GetUserDataSession from session
func GetUserDataSession(ctx context.Context) uint {
	if _, ok := ctx.Value(config.ContextKeyUser).(uint); !ok {
		return 0
	}

	return ctx.Value(config.ContextKeyUser).(uint)
}
