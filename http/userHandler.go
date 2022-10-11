package http

import (
	"encoding/json"
	"net/http"
	"vending-machine/app"
	"vending-machine/domain"
)

type CreateUserHandler struct {
	service app.UserCrudService
}

func NewCreateUserHandler(s app.UserCrudService) *CreateUserHandler {
	return &CreateUserHandler{
		service: s,
	}
}

func (h *CreateUserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
