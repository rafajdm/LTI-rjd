package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rafajdm/lti-rjd/application/usecases"
	"github.com/rafajdm/lti-rjd/domain/models"
)

type UserController struct {
	CreateCandidateProfileUseCase usecases.CreateCandidateProfileUseCaseInterface
}

func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.CreateCandidateProfileUseCase.Execute(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
