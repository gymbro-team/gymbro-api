package handler

import (
	"gymbro-api/controller"

	"github.com/gorilla/mux"
)

type ExerciseHandler struct {
	workoutController *controller.ExerciseController
}

func NewExerciseHandler(workoutController *controller.ExerciseController) *ExerciseHandler {
	return &ExerciseHandler{workoutController}

}
func (wh *ExerciseHandler) RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/exercises", wh.workoutController.CreateExerciseHandler).Methods("POST")
	mux.HandleFunc("/exercises", wh.workoutController.GetExercisesHandler).Methods("GET")
	mux.HandleFunc("/exercises/{id}", wh.workoutController.GetExerciseHandler).Methods("GET")
	mux.HandleFunc("/exercises/{id}", wh.workoutController.UpdateExerciseHandler).Methods("PUT")
	mux.HandleFunc("/exercises/{id}", wh.workoutController.DeleteExerciseHandler).Methods("DELETE")
}
