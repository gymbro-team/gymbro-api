package controller

import (
	"encoding/json"
	"gymbro-api/model"
	"gymbro-api/repository"
	"net/http"
)

type SessionController struct {
	sessionRepo *repository.SessionRepository
}

func NewSessionController(sessionRepo *repository.SessionRepository) *SessionController {
	return &SessionController{sessionRepo}
}

func (sc *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	var login model.Login

	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		http.Error(w, "Failed to decode request body, verify your data type and fields", http.StatusBadRequest)
		return
	}

	session, err := sc.sessionRepo.Login(&login)

	if err != nil {
		http.Error(w, "Failed to login", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(session)
}
