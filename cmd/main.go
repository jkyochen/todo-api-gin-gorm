package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"todo-api-gin-gorm/pkg/models"
	"todo-api-gin-gorm/pkg/routers"
)

func main() {

	if err := models.NewEngine(); err != nil {
		log.Fatal("Failed to initialize ORM engine:", err)
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      routers.Load(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
