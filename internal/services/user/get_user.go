package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	user.Email = email

	err := database.DB.QueryRow("SELECT id, first_name, last_name, password, phone_number, college, gender, role, country, github, bio, is_banned, is_added, is_vitian, is_verified, team_id FROM users WHERE email = $1",
		email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Phone,
		&user.College, &user.Gender, &user.Role, &user.Country, &user.Github, &user.Bio,
		&user.IsBanned, &user.IsAdded, &user.IsVitian, &user.IsVerified, &user.TeamID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
