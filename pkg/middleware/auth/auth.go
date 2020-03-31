package auth

import (
	"context"
	"net/http"

	"todo-api-gin-gorm/pkg/helper"
	"todo-api-gin-gorm/pkg/models"

	"github.com/gin-gonic/gin"
)

// Check initializes the auth middleware.
func Check() gin.HandlerFunc {

	return func(c *gin.Context) {

		session, err := models.SessionStore.Get(c.Request, "session")
		if err != nil {
			helper.ErrorJSON(c, http.StatusUnauthorized, helper.ErrNotLogin)
			return
		}

		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			helper.ErrorJSON(c, http.StatusUnauthorized, helper.ErrNotLogin)
			return
		}

		ctx := context.WithValue(c.Request.Context(), "user_id", userID)
		c.Request = c.Request.WithContext(ctx)
	}
}
