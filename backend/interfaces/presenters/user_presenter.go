package presenters

import "github.com/rafajdm/lti-rjd/domain/models"

type UserPresenter struct{}

func (p *UserPresenter) Present(user *models.User) map[string]interface{} {
	// Convert user model to a map or any other presentation format
	return map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email, // Ensure this line is correct
	}
}
