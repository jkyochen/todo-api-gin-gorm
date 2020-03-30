package routers

import (
	"net/http"

	"todo-api-gin-gorm/api"

	"github.com/gin-gonic/gin"
)

// Load initializes the routing of the application.
func Load(middleware ...gin.HandlerFunc) http.Handler {

	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(middleware...)

	// 404 not found
	e.NoRoute(api.NotFound)

	e.GET("/healthz", api.Heartbeat)
	return e
}
