package test

import (
	"net/http"
	"testing"

	"todo-api-gin-gorm/pkg/models"
	"todo-api-gin-gorm/pkg/routers"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	assert.NoError(t, models.PrepareTestDatabase())

	r := gofight.New()

	t.Run("success register", func(t *testing.T) {
		r.POST("/register").
			SetJSON(gofight.D{
				"name":     "jack",
				"password": "jack",
				"email":    "jack@mail.com",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				id, _ := jsonparser.GetInt(data, "id")
				created, _ := jsonparser.GetString(data, "created_at")
				assert.Equal(t, http.StatusOK, r.Code)
				assert.NotNil(t, id)
				assert.NotNil(t, created)
			})
	})

	t.Run("register don't input", func(t *testing.T) {
		r.POST("/register").
			SetJSON(gofight.D{}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusBadRequest, r.Code)
				assert.Equal(t, 101, int(code))
				assert.Equal(t, "Error Code: 101, Error Message: Bad Input Request", err)
			})
	})

	t.Run("duplicate register", func(t *testing.T) {
		r.POST("/register").
			SetJSON(gofight.D{
				"name":     "jack",
				"password": "jack",
				"email":    "jack@mail.com",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusConflict, r.Code)
				assert.Equal(t, 103, int(code))
				assert.Equal(t, "Error Code: 103, Error Message: User Exist", err)
			})
	})
}
