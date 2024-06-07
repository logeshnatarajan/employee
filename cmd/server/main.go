package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/logeshnatarajan/employee/config"
	"github.com/logeshnatarajan/employee/internal/handler"
	"github.com/logeshnatarajan/employee/internal/repository"
	"github.com/logeshnatarajan/employee/internal/service"
	"github.com/logeshnatarajan/employee/internal/store"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the store
	db, err := store.NewPostgresStore(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Set up the router
	r := gin.Default()

	repo := repository.NewPostgresRepository(db)
	service := service.NewEmployeeService(repo)
	handler.RegisterRoutes(r, service)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
