package test

import (
	"net/http"
	"testing"

	"todo-api-gin-gorm/pkg/routers"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	r := gofight.New()

	t.Run("return 200", func(t *testing.T) {
		r.GET("/healthz").
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				assert.Equal(t, "ok", r.Body.String())
			})
	})
}
