package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	user.Email = email

	var check uuid.NullUUID

	err := database.DB.QueryRow("SELECT id, first_name, last_name, reg_no, password, phone, college, gender, role, country, github, bio, is_banned, is_added, is_vitian, is_verified, team_id FROM users WHERE email = $1",
		email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.RegNo, &user.Password, &user.Phone,
		&user.College, &user.Gender, &user.Role, &user.Country, &user.Github, &user.Bio,
		&user.IsBanned, &user.IsAdded, &user.IsVitian, &user.IsVerified, &check)
	if err != nil {
		return nil, err
	}

	if check.Valid {
		user.TeamID = check.UUID
	} else {
		user.TeamID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
	}

	return &user, nil
}
