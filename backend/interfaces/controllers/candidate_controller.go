package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rafajdm/lti-rjd/application/usecases"
	"github.com/rafajdm/lti-rjd/domain/models"
)

type CandidateController struct {
	CreateCandidateProfileUseCase usecases.CreateCandidateProfileUseCaseInterface
}

func (c *CandidateController) CreateCandidateHandler(w http.ResponseWriter, r *http.Request) {
	var candidate struct {
		Name            string            `json:"name"`
		ContactInfo     map[string]string `json:"contact_info"`
		ResumeLink      string            `json:"resume_link"`
		CurrentPosition string            `json:"current_position"`
		Location        string            `json:"location"`
	}

	// Read the raw request body for debugging
	body, _ := json.Marshal(candidate)
	log.Printf("Raw request body: %s", string(body))

	if err := json.NewDecoder(r.Body).Decode(&candidate); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received candidate data: %+v", candidate) // Log the received data

	user := models.User{
		Name:            candidate.Name,
		ContactInfo:     candidate.ContactInfo,
		ResumeLink:      candidate.ResumeLink,
		CurrentPosition: candidate.CurrentPosition,
		Location:        candidate.Location,
	}

	log.Printf("Created user model: %+v", user)

	if err := c.CreateCandidateProfileUseCase.Execute(&user); err != nil {
		log.Printf("Error executing use case: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
