package persistence

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/rafajdm/lti-rjd/domain/models"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func (r *UserRepositoryImpl) Save(user *models.User) error {
	contactInfoJSON, err := json.Marshal(user.ContactInfo)
	if err != nil {
		return err
	}

	log.Printf("Inserting user data: Name=%s, ContactInfo=%s, ResumeLink=%s, CurrentPosition=%s, Location=%s",
		user.Name, contactInfoJSON, user.ResumeLink, user.CurrentPosition, user.Location)

	query := `
        INSERT INTO CandidateProfile (name, contact_info, resume_link, current_position, location)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err = r.DB.Exec(query, user.Name, contactInfoJSON, user.ResumeLink, user.CurrentPosition, user.Location)
	return err
}

func (r *UserRepositoryImpl) FindByID(id int) (*models.User, error) {
	query := `
        SELECT candidate_id, name, contact_info, resume_link, current_position, location
        FROM CandidateProfile
        WHERE candidate_id = $1
    `
	row := r.DB.QueryRow(query, id)

	var user models.User
	var contactInfoJSON string
	err := row.Scan(&user.ID, &user.Name, &contactInfoJSON, &user.ResumeLink, &user.CurrentPosition, &user.Location)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(contactInfoJSON), &user.ContactInfo)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
