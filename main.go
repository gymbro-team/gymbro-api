package main

import (
	"gymbro-api/config"
	"gymbro-api/controller"
	"gymbro-api/core"
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

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NewApp(cfg *config.Config) *App {
	core.InitializeDatabase("user=" + cfg.Database.User + " password=" + cfg.Database.Password + " dbname=" + cfg.Database.Name + " sslmode=disable")

	app := &App{
		Router: mux.NewRouter(),
	}
	app.Router.Use(jsonMiddleware)

	v1 := app.Router.PathPrefix("/v1").Subrouter()
	auth := v1.PathPrefix("/auth").Subrouter()

	userRepo := repository.NewUserRepository(core.GetDB())
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)
	userHandler.RegisterRoutes(v1)

	sessionRepo := repository.NewSessionRepository(core.GetDB())
	sessionController := controller.NewSessionController(sessionRepo)
	sessionHandler := handler.NewSessionHandler(sessionController)
	sessionHandler.RegisterRoutes(auth)

	workoutRepo := repository.NewWorkoutRepository(core.GetDB())
	workoutController := controller.NewWorkoutController(workoutRepo)
	workoutHandler := handler.NewWorkoutHandler(workoutController)
	workoutHandler.RegisterRoutes(v1)

	return app
}

func main() {
	cfg := config.LoadConfig()

	app := NewApp(cfg)

	addr := cfg.ServerAddress

	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, app.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
