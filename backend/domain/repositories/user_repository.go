package repositories

import "github.com/rafajdm/lti-rjd/domain/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByID(id int) (*models.User, error)
}
