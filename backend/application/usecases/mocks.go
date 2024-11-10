package usecases

import (
	"github.com/rafajdm/lti-rjd/domain/models"
)

type MockUserRepository struct {
	SaveFunc     func(user *models.User) error
	FindByIDFunc func(id int) (*models.User, error)
}

func (m *MockUserRepository) Save(user *models.User) error {
	return m.SaveFunc(user)
}

func (m *MockUserRepository) FindByID(id int) (*models.User, error) {
	return m.FindByIDFunc(id)
}
