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
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", "user="+cfg.Database.User+" password="+cfg.Database.Password+" dbname="+cfg.Database.Name+" sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)

	mux := mux.NewRouter()

	userHandler.RegisterRoutes(mux)

	addr := cfg.ServerAddress
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
