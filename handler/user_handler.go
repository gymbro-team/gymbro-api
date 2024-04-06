package handler

import (
	"gymbro-api/controller"
	"net/http"
)

type UserHandler struct {
	userController *controller.UserController
}

func NewUserHandler(userController *controller.UserController) *UserHandler {
	return &UserHandler{userController}
}

func (uh *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	// Define routes for user CRUD operations
	mux.HandleFunc("/users", uh.userController.CreateUserHandler)
	mux.HandleFunc("/users/{id}", uh.userController.GetUserHandler)
}
