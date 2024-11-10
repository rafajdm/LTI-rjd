package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafajdm/lti-rjd/domain/models"
	"github.com/stretchr/testify/assert"
)

type MockCreateCandidateProfileUseCase struct {
	ExecuteFunc func(user *models.User) error
}

func (m *MockCreateCandidateProfileUseCase) Execute(user *models.User) error {
	return m.ExecuteFunc(user)
}

func TestUserController_CreateUserHandler(t *testing.T) {
	mockUseCase := &MockCreateCandidateProfileUseCase{
		ExecuteFunc: func(user *models.User) error {
			return nil
		},
	}

	controller := &UserController{
		CreateCandidateProfileUseCase: mockUseCase,
	}

	user := &models.User{
		Name:            "John Doe",
		ContactInfo:     `{"email": "john.doe@example.com", "phone": "123-456-7890"}`,
		ResumeLink:      "http://example.com/resume.pdf",
		CurrentPosition: "Software Engineer",
		Location:        "New York",
	}

	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateUserHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status Created")

	// Check the response body is what we expect.
	expected := ``
	assert.Equal(t, expected, rr.Body.String(), "Expected response body to be empty")
}
