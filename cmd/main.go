package main

import (
	"net/http"
	"time"

	"todo-api-gin-gorm/pkg/routers"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      routers.Load(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
