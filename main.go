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

type App struct {
	Router *mux.Router
}

func NewApp(db *sql.DB) *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	v1 := app.Router.PathPrefix("/v1").Subrouter()

	public := v1.PathPrefix("/public").Subrouter()

	userRepo := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)
	userHandler.RegisterRoutes(v1)

	sessionRepo := repository.NewSessionRepository(db)
	sessionController := controller.NewSessionController(sessionRepo)
	sessionHandler := handler.NewSessionHandler(sessionController)
	sessionHandler.RegisterRoutes(public)

	return app
}

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", "user="+cfg.Database.User+" password="+cfg.Database.Password+" dbname="+cfg.Database.Name+" sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	app := NewApp(db)

	addr := cfg.ServerAddress
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, app.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
