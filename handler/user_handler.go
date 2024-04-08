package handler

import (
	"gymbro-api/controller"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userController *controller.UserController
}

func NewUserHandler(userController *controller.UserController) *UserHandler {
	return &UserHandler{userController}
}

func (uh *UserHandler) RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/users", uh.userController.CreateUserHandler).Methods("POST")
	mux.HandleFunc("/users", uh.userController.GetUsersHandler).Methods("GET")
	mux.HandleFunc("/users/{id}", uh.userController.GetUserHandler).Methods("GET")
	mux.HandleFunc("/users/{id}", uh.userController.UpdateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id}", uh.userController.DeleteUserHandler).Methods("DELETE")
}
