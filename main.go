package main

import (
	"database/sql"
	"gymbro-api/config"
	"gymbro-api/controller"
	"gymbro-api/handler"
	"gymbro-api/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := sql.Open("postgres", "user="+cfg.Database.User+" password="+cfg.Database.Password+" dbname="+cfg.Database.Name+" sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create repository instances
	userRepo := repository.NewUserRepository(db)

	// Create controller instances
	userController := controller.NewUserController(userRepo)

	// Create handler instances
	userHandler := handler.NewUserHandler(userController)

	// Create a ServeMux (router)
	mux := mux.NewRouter()

	// Register user-related routes
	userHandler.RegisterRoutes(mux)

	// Start the HTTP server
	addr := cfg.ServerAddress
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
