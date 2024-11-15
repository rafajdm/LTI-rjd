package usecases

import (
	"github.com/rafajdm/lti-rjd/domain/models"
	"github.com/rafajdm/lti-rjd/domain/repositories"
)

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func (uc *CreateUserUseCase) Execute(user *models.User) error {
	return uc.UserRepository.Save(user)
}
