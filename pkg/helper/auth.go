package helper

import "context"

// GetUserDataSession from session
func GetUserDataSession(ctx context.Context) uint {
	if _, ok := ctx.Value("user_id").(uint); !ok {
		return 0
	}

	return ctx.Value("user_id").(uint)
}
