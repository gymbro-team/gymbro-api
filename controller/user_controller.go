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

type UserController struct {
	userRepo *repository.UserRepository
}

func NewUserController(userRepo *repository.UserRepository) *UserController {
	return &UserController{userRepo}
}

func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Failed to decode request body, verify your data type and fields", http.StatusBadRequest)
		return
	}

	err = uc.userRepo.CreateUser(&user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uc.userRepo.GetUserByID(id)

	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userRepo.GetUsers()

	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Failed to decode request body, verify your data type and fields", http.StatusBadRequest)
		return
	}

	user.ID = id

	err = uc.userRepo.UpdateUser(&user)

	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = uc.userRepo.DeleteUser(id)

	if err != nil {
		if err == repository.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
