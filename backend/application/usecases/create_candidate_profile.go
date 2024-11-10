package usecases

import (
	"github.com/rafajdm/lti-rjd/domain/models"
	"github.com/rafajdm/lti-rjd/domain/repositories"
)

type CreateCandidateProfileUseCase struct {
	UserRepository repositories.UserRepository
}

func (uc *CreateCandidateProfileUseCase) Execute(user *models.User) error {
	return uc.UserRepository.Save(user)
}
