package usecases

import "github.com/rafajdm/lti-rjd/domain/models"

type CreateCandidateProfileUseCaseInterface interface {
	Execute(user *models.User) error
}
