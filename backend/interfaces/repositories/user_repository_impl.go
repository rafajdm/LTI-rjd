package repositories

import (
	"database/sql"

	"github.com/rafajdm/lti-rjd/domain/models"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func (r *UserRepositoryImpl) Save(user *models.User) error {
	// Implementation for saving a user to the database
	return nil
}

func (r *UserRepositoryImpl) FindByID(id int) (*models.User, error) {
	// Implementation for finding a user by ID from the database
	return nil, nil
}
