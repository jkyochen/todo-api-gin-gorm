package routers

import (
	"net/http"

	"todo-api-gin-gorm/api"
	"todo-api-gin-gorm/pkg/config"
	"todo-api-gin-gorm/pkg/middleware/auth"
	"todo-api-gin-gorm/pkg/middleware/header"

	"github.com/gin-gonic/gin"
)

// Load initializes the routing of the application.
func Load(middleware ...gin.HandlerFunc) http.Handler {

	if config.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())

	e.Use(header.Options)
	e.Use(header.Secure)
	e.Use(middleware...)

	root := e.Group(config.Server.Root)
	{
		// user
		root.POST("/register", api.Register)
		root.POST("/login", api.Login)

		// todo
		v := root.Group("")
		v.Use(auth.Check())
		{
			v.POST("/todo", api.CreateTodo)
			v.PUT("/todo/:id", api.UpdateTodo)
			v.DELETE("/todo/:id", api.DeleteTodo)
			v.GET("/todo/:id", api.GetTodo)
			v.GET("/todo", api.GetTodos)
		}

		root.GET("/healthz", api.Heartbeat)
	}

	// 404 not found
	e.NoRoute(api.NotFound)

	return e
}
