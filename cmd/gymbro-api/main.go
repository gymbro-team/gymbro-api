package main

import (
	"gymbro-api/auth"
	"gymbro-api/config"
	"gymbro-api/controller"
	"gymbro-api/core"
	"gymbro-api/handler"
	"gymbro-api/repository"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

	authRouter := app.Router.PathPrefix("/auth/v1").Subrouter()
	v1 := app.Router.PathPrefix("/api/v1").Subrouter()

	v1.Use(auth.AuthecationMiddleware)

	userRepo := repository.NewUserRepository(core.GetDB())
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)
	userHandler.RegisterRoutes(v1)

	sessionRepo := repository.NewSessionRepository(core.GetDB())
	sessionController := controller.NewSessionController(sessionRepo)
	sessionHandler := handler.NewSessionHandler(sessionController)
	sessionHandler.RegisterRoutes(authRouter)

	workoutRepo := repository.NewWorkoutRepository(core.GetDB())
	workoutController := controller.NewWorkoutController(workoutRepo)
	workoutHandler := handler.NewWorkoutHandler(workoutController)
	workoutHandler.RegisterRoutes(v1)

	exerciseRepo := repository.NewExerciseRepository(core.GetDB())
	exerciseController := controller.NewExerciseController(exerciseRepo)
	exerciseHandler := handler.NewExerciseHandler(exerciseController)
	exerciseHandler.RegisterRoutes(v1)

	return app
}

func main() {
	cfg := config.LoadConfig()

	app := NewApp(cfg)

	addr := cfg.ServerAddress
	// CORS options
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(app.Router)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
