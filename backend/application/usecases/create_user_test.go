package usecases

import (
	"testing"

	"github.com/rafajdm/lti-rjd/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	mockRepo := &MockUserRepository{
		SaveFunc: func(user *models.User) error {
			return nil
		},
	}

	useCase := &CreateUserUseCase{
		UserRepository: mockRepo,
	}

	user := &models.User{
		Name:            "John Doe",
		ContactInfo:     `{"email": "john.doe@example.com", "phone": "123-456-7890"}`,
		ResumeLink:      "http://example.com/resume.pdf",
		CurrentPosition: "Software Engineer",
		Location:        "New York",
	}

	err := useCase.Execute(user)
	assert.NoError(t, err)
}
