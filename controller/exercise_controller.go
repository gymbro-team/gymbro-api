package controller

import (
	"encoding/json"
	"gymbro-api/auth"
	"gymbro-api/model"
	"gymbro-api/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ExerciseController struct {
	exerciseRepo *repository.ExerciseRepository
}

func NewExerciseController(exerciseRepo *repository.ExerciseRepository) *ExerciseController {
	return &ExerciseController{exerciseRepo}
}

func (uc *ExerciseController) CreateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var exercise model.Exercise

	err := json.NewDecoder(r.Body).Decode(&exercise)

	if err != nil {
		http.Error(w, "Failed to decode request body, verify your data type and fields", http.StatusBadRequest)
		return
	}

	userId, err := auth.GetParsedUserId(r.Header.Get("user_id"))

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = uc.exerciseRepo.CreateExercise(&exercise, userId)
	if err != nil {
		log.Printf("Failed to create exercise: %v", err)
		http.Error(w, "Failed to create exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *ExerciseController) GetExerciseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	exercise, err := uc.exerciseRepo.GetExerciseByID(id)

	if err != nil {
		if err == repository.ErrExerciseNotFound {
			http.Error(w, "Exercise not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get exercise", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exercise)
}

func (uc *ExerciseController) GetExercisesHandler(w http.ResponseWriter, r *http.Request) {
	exercises, err := uc.exerciseRepo.GetExercises()

	if len(exercises) == 0 {
		http.Error(w, "No exercises found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Failed to get exercises", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exercises)
}

func (uc *ExerciseController) UpdateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	var exercise model.Exercise

	err = json.NewDecoder(r.Body).Decode(&exercise)

	if err != nil {
		http.Error(w, "Failed to decode request body, verify your data type and fields", http.StatusBadRequest)
		return
	}

	exercise.ID = id

	err = uc.exerciseRepo.UpdateExercise(&exercise)

	if err != nil {
		if err == repository.ErrExerciseNotFound {
			http.Error(w, "Exercise not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *ExerciseController) DeleteExerciseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	err = uc.exerciseRepo.DeleteExercise(id)

	if err != nil {
		if err == repository.ErrExerciseNotFound {
			http.Error(w, "Exercise not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
