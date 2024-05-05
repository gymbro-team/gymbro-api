package handler

import (
	"gymbro-api/controller"

	"github.com/gorilla/mux"
)

type WorkoutHandler struct {
	workoutController *controller.WorkoutController
}

func NewWorkoutHandler(workoutController *controller.WorkoutController) *WorkoutHandler {
	return &WorkoutHandler{workoutController}

}
func (wh *WorkoutHandler) RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/workouts", wh.workoutController.CreateWorkoutHandler).Methods("POST")
	mux.HandleFunc("/workouts", wh.workoutController.GetWorkoutsHandler).Methods("GET")
	mux.HandleFunc("/workouts/{id}", wh.workoutController.GetWorkoutHandler).Methods("GET")
	mux.HandleFunc("/workouts/{id}", wh.workoutController.UpdateWorkoutHandler).Methods("PUT")
	mux.HandleFunc("/workouts/{id}", wh.workoutController.DeleteWorkoutHandler).Methods("DELETE")
}
