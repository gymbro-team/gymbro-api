package controller

import (
	"encoding/json"
	"gymbro-api/model"
	"gymbro-api/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type WorkoutController struct {
	workoutRepo *repository.WorkoutRepository
}

func NewWorkoutController(workoutRepo *repository.WorkoutRepository) *WorkoutController {
	return &WorkoutController{workoutRepo}
}

func (uc *WorkoutController) CreateWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	var workout model.Workout

	err := json.NewDecoder(r.Body).Decode(&workout)

	if err != nil {
		http.Error(w, "Failed to decode request body, verify your data type and fields", http.StatusBadRequest)
		return
	}

	err = uc.workoutRepo.CreateWorkout(&workout)
	if err != nil {
		log.Printf("Failed to create workout: %v", err)
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *WorkoutController) GetWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	workout, err := uc.workoutRepo.GetWorkoutByID(id)

	if err != nil {
		if err == repository.ErrWorkoutNotFound {
			http.Error(w, "Workout not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get workout", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(workout)
}

func (uc *WorkoutController) GetWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	workouts, err := uc.workoutRepo.GetWorkouts()

	if len(workouts) == 0 {
		http.Error(w, "No workouts found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Failed to get workouts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(workouts)
}

func (uc *WorkoutController) UpdateWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	var workout model.Workout

	err = json.NewDecoder(r.Body).Decode(&workout)

	if err != nil {
		http.Error(w, "Failed to decode request body, verify your data type and fields", http.StatusBadRequest)
		return
	}

	workout.ID = id

	err = uc.workoutRepo.UpdateWorkout(&workout)

	if err != nil {
		if err == repository.ErrWorkoutNotFound {
			http.Error(w, "Workout not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *WorkoutController) DeleteWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	err = uc.workoutRepo.DeleteWorkout(id)

	if err != nil {
		if err == repository.ErrWorkoutNotFound {
			http.Error(w, "Workout not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
