package test

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"todo-api-gin-gorm/pkg/models"
	"todo-api-gin-gorm/pkg/routers"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	assert.NoError(t, models.PrepareTestDatabase())

	t.Run("success create", func(t *testing.T) {

		r := gofight.New()

		var cookie string
		r.POST("/login").
			SetJSON(gofight.D{
				"email":    "check@mail.com",
				"password": "check",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				cookie = r.HeaderMap.Get("Set-Cookie")
				assert.NotEmpty(t, cookie)
			})

		var id int64

		r.POST("/todo").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			SetJSON(gofight.D{
				"title": "Study English",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				id, _ = jsonparser.GetInt(data, "id")
				created, _ := jsonparser.GetString(data, "created_at")
				assert.Equal(t, http.StatusOK, r.Code)
				assert.NotNil(t, id)
				assert.NotNil(t, created)
			})

		r.GET("/todo/"+strconv.FormatInt(id, 10)).
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				title, _ := jsonparser.GetString(data, "Title")
				content, _ := jsonparser.GetString(data, "Content")
				assert.Equal(t, http.StatusOK, r.Code)
				assert.Equal(t, "Study English", title)
				assert.Equal(t, "", content)
			})
	})

	t.Run("create don't login", func(t *testing.T) {

		r := gofight.New()

		r.POST("/todo").
			SetJSON(gofight.D{
				"title": "Study English",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusUnauthorized, r.Code)
				assert.Equal(t, 102, int(code))
				assert.Equal(t, "Error Code: 102, Error Message: User Not Login", err)
			})
	})

	t.Run("create don't input", func(t *testing.T) {

		r := gofight.New()

		var cookie string
		r.POST("/login").
			SetJSON(gofight.D{
				"email":    "check@mail.com",
				"password": "check",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				cookie = r.HeaderMap.Get("Set-Cookie")
				assert.NotEmpty(t, cookie)
			})

		r.POST("/todo").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
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
}

func TestUpdateTodo(t *testing.T) {
	assert.NoError(t, models.PrepareTestDatabase())

	t.Run("success update", func(t *testing.T) {

		r := gofight.New()

		var cookie string
		r.POST("/login").
			SetJSON(gofight.D{
				"email":    "check@mail.com",
				"password": "check",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				cookie = r.HeaderMap.Get("Set-Cookie")
				assert.NotEmpty(t, cookie)
			})

		r.PUT("/todo/1").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			SetJSON(gofight.D{
				"title": "Study English",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				id, _ := jsonparser.GetInt(data, "id")
				updated, _ := jsonparser.GetString(data, "updated_at")
				assert.Equal(t, http.StatusOK, r.Code)
				assert.NotNil(t, id)
				assert.NotNil(t, updated)
			})

		r.GET("/todo/1").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				title, _ := jsonparser.GetString(data, "Title")
				content, _ := jsonparser.GetString(data, "Content")
				assert.Equal(t, http.StatusOK, r.Code)
				assert.Equal(t, "Study English", title)
				assert.Equal(t, "", content)
			})
	})

	t.Run("update don't login", func(t *testing.T) {

		r := gofight.New()

		r.PUT("/todo/1").
			SetJSON(gofight.D{
				"title": "Study English",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusUnauthorized, r.Code)
				assert.Equal(t, 102, int(code))
				assert.Equal(t, "Error Code: 102, Error Message: User Not Login", err)
			})
	})

	t.Run("update don't input", func(t *testing.T) {

		r := gofight.New()

		var cookie string
		r.POST("/login").
			SetJSON(gofight.D{
				"email":    "check@mail.com",
				"password": "check",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				cookie = r.HeaderMap.Get("Set-Cookie")
				assert.NotEmpty(t, cookie)
			})

		r.PUT("/todo/1").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
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
}

func TestDeleteTodo(t *testing.T) {
	assert.NoError(t, models.PrepareTestDatabase())

	t.Run("success delete", func(t *testing.T) {

		r := gofight.New()

		var cookie string
		r.POST("/login").
			SetJSON(gofight.D{
				"email":    "check@mail.com",
				"password": "check",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				cookie = r.HeaderMap.Get("Set-Cookie")
				assert.NotEmpty(t, cookie)
			})

		r.DELETE("/todo/1").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				id, _ := jsonparser.GetInt(data, "id")
				updated, _ := jsonparser.GetString(data, "updated_at")
				assert.Equal(t, http.StatusOK, r.Code)
				assert.NotNil(t, id)
				assert.NotNil(t, updated)
			})

		r.GET("/todo/1").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusForbidden, r.Code)
				assert.Equal(t, 104, int(code))
				assert.Equal(t, "Error Code: 104, Error Message: Todo Not Exist", err)
			})
	})

	t.Run("delete don't login", func(t *testing.T) {

		r := gofight.New()

		r.DELETE("/todo/1").
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusUnauthorized, r.Code)
				assert.Equal(t, 102, int(code))
				assert.Equal(t, "Error Code: 102, Error Message: User Not Login", err)
			})
	})
}

func TestGetTodo(t *testing.T) {
	assert.NoError(t, models.PrepareTestDatabase())

	t.Run("success get", func(t *testing.T) {

		r := gofight.New()

		var cookie string
		r.POST("/login").
			SetJSON(gofight.D{
				"email":    "check@mail.com",
				"password": "check",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				cookie = r.HeaderMap.Get("Set-Cookie")
				assert.NotEmpty(t, cookie)
			})

		r.GET("/todo/1").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				title, _ := jsonparser.GetString(data, "Title")
				content, _ := jsonparser.GetString(data, "Content")
				assert.Equal(t, http.StatusOK, r.Code)
				assert.Equal(t, "学习数学", title)
				assert.Equal(t, "学习第一章", content)
			})
	})

	t.Run("get don't login", func(t *testing.T) {

		r := gofight.New()

		r.GET("/todo/1").
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusUnauthorized, r.Code)
				assert.Equal(t, 102, int(code))
				assert.Equal(t, "Error Code: 102, Error Message: User Not Login", err)
			})
	})
}

func TestGetTodos(t *testing.T) {
	assert.NoError(t, models.PrepareTestDatabase())

	t.Run("success get", func(t *testing.T) {

		r := gofight.New()

		var cookie string
		r.POST("/login").
			SetJSON(gofight.D{
				"email":    "check@mail.com",
				"password": "check",
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				cookie = r.HeaderMap.Get("Set-Cookie")
				assert.NotEmpty(t, cookie)
			})

		r.GET("/todo").
			SetHeader(gofight.H{
				"Cookie": cookie,
			}).
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				fmt.Println("data", string(data))
				todos := make([]string, 0, 1)
				jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					v, _ := jsonparser.GetString(value, "value")
					todos = append(todos, v)
				})
				assert.Equal(t, http.StatusOK, r.Code)
				assert.Equal(t, 1, len(todos))
			})
	})

	t.Run("get don't login", func(t *testing.T) {

		r := gofight.New()

		r.GET("/todo").
			Run(routers.Load(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				data := r.Body.Bytes()
				code, _ := jsonparser.GetInt(data, "code")
				err, _ := jsonparser.GetString(data, "error")
				assert.Equal(t, http.StatusUnauthorized, r.Code)
				assert.Equal(t, 102, int(code))
				assert.Equal(t, "Error Code: 102, Error Message: User Not Login", err)
			})
	})
}
