package routers

import (
	"net/http"

	"todo-api-gin-gorm/api"
	"todo-api-gin-gorm/pkg/middleware/auth"
	"todo-api-gin-gorm/pkg/middleware/header"

	"github.com/gin-gonic/gin"
)

// Load initializes the routing of the application.
func Load(middleware ...gin.HandlerFunc) http.Handler {

	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(gin.Logger())

	e.Use(header.Options)
	e.Use(header.Secure)
	e.Use(middleware...)

	// user
	e.POST("/register", api.Register)
	e.POST("/login", api.Login)

	// todo
	v := e.Group("")
	v.Use(auth.Check())
	{
		v.POST("/todo", api.CreateTodo)
		v.PUT("/todo/:id", api.UpdateTodo)
		v.DELETE("/todo/:id", api.DeleteTodo)
		v.GET("/todo/:id", api.GetTodo)
		v.GET("/todo", api.GetTodos)
	}

	// 404 not found
	e.NoRoute(api.NotFound)

	e.GET("/healthz", api.Heartbeat)
	return e
}
